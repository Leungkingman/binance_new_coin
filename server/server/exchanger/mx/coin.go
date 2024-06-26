package mx

import (
	"binanceNewCoin/server/global"
	"binanceNewCoin/server/logger"
	"binanceNewCoin/server/model"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func GetAllCoin() ([]string, error) {
	var coins = []string{}
	completeUrl := global.MX_HOST + "/open/api/v2/market/symbols"
	resp, err := http.Get(completeUrl)
	defer resp.Body.Close()
	if err != nil {
		//fmt.Printf("err = ", err.Error())
		return coins, err
	} else {
		result, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return coins, err
		}
		//var mxSymbolResult model.MxSymbolResult
		//err = json.Unmarshal(result, &mxSymbolResult)
		//fmt.Printf("result = ", string(result))
		symbols := gjson.Get(string(result), "data.#.symbol").Array()
		//fmt.Printf("mx symbols = ", symbols)
		for _, item := range symbols {
			if strings.Contains(item.Str, "_USDT") {
				coin := item.Str[0:strings.Index(item.Str, "_USDT")]
				coins = append(coins, coin)
			}
		}
	}
	return coins, nil
}

// 维护 global.MX_MONITOR_COIN 的币种价格池
func SetMxCoinPrice() {

	global.MX_MONITOR_UPDATE_TIMES = 0
	coinPriceChannel := make(chan model.CoinPrice)
	coinErrorChannel := make(chan model.CoinError)

	// 本地开发环境使用
	//for i := 0; i < 2; i ++ {
	//	go func(i int) {
	//		GetMxCoinPrice(coinPriceChannel, coinErrorChannel, global.MX_MONITOR_COIN_ARRAY[i * 200: (i+1) * 200])
	//	}(i)
	//}
	//go func() {
	//	GetMxCoinPrice(coinPriceChannel, coinErrorChannel, global.MX_MONITOR_COIN_ARRAY[400: len(global.MX_MONITOR_COIN_ARRAY)])
	//}()

	// 正式环境使用
	go func() {
		GetMxCoinPrice(coinPriceChannel, coinErrorChannel, global.MX_MONITOR_COIN_ARRAY)
	}()

	coinPrice := model.CoinPrice{}
	coinError := model.CoinError{}

	logger.ClearNormmalLog("exchange")
	getTimes := 0

	//fmt.Printf("准备开始，总次数为：%v \n", len(global.MX_MONITOR_COIN_ARRAY))

	for {
		//fmt.Printf("getTimes = %v \n", getTimes)
		if getTimes == len(global.MX_MONITOR_COIN_ARRAY) {
			fmt.Printf("MX_MONITOR_COIN = %v \n", global.MX_MONITOR_COIN)
			fmt.Printf("初始化已完成 \n")
			global.INIT_EXCHANGE_FINISH = true
			global.LAST_UPDATE_MX_MONITOR_TIME = time.Now().Unix()
			global.LAST_UPDATE_MX_MONITOR_AMOUNT = global.MX_MONITOR_UPDATE_TIMES
			global.MX_MONITOR_UPDATING = false
			close(coinPriceChannel)
			close(coinErrorChannel)
			break
		}
		select {
			case coinPrice, _ = <- coinPriceChannel:
				getTimes += 1
				global.MX_MONITOR_UPDATE_TIMES += 1
				priceStr := coinPrice.Price
				prec := 0
				if strings.Contains(priceStr, ".") {
					startIndex := strings.Index(priceStr, ".")
					priceLength := len(priceStr)
					prec = len(priceStr[startIndex + 1: priceLength])
				}
				price, _ := strconv.ParseFloat(priceStr,64)
				global.MX_MONITOR_COIN[coinPrice.Coin] = model.CoinItem{
					Price: price,
					Prec: prec,
				}
				content := coinPrice.Coin + " : " + coinPrice.Price
				logger.Addlog(content, "normal", "exchange")
			case coinError, _ = <- coinErrorChannel:
				//fmt.Printf("coin = %v \n ", coinError.Coin)
				getTimes += 1
				if coinError.Err != nil {
					content := coinError.Coin + " : " + coinError.Err.Error()
					//fmt.Printf("contentr = %v \n ", content)
					logger.Addlog(content, "error", "error")
				}
		}
	}
}

// 维护 global.MX_MONITOR_COIN 的币种价格池
func GetMxCoinPrice(cpChannel chan model.CoinPrice, errChannel chan model.CoinError, coins []string) {
	//fmt.Printf("coins = %v \n", coins)
	//fmt.Printf("============================================== \n")
	for _, v := range coins {
		//fmt.Printf("coin = %v \n ", v)
		url := global.MX_HOST + "/open/api/v2/market/ticker?symbol=" + v + "_USDT"
		//fmt.Printf("url = %v \n", url)
		resp, httpErr := http.Get(url)
		if httpErr != nil {
			//fmt.Printf("httpErr = %v \n", httpErr.Error())
			coinErr := model.CoinError{
				Coin: v,
				Err: httpErr,
			}
			errChannel <- coinErr
		}
		defer resp.Body.Close()
		result, readErr := ioutil.ReadAll(resp.Body)
		if readErr != nil {
			coinErr := model.CoinError{
				Coin: v,
				Err: readErr,
			}
			errChannel <- coinErr
		}
		var tickerResult GetResult
		_ = json.Unmarshal(result, &tickerResult)
		if tickerResult.Code == 200 && len(tickerResult.Data) > 0 {
			//global.MX_MONITOR_COIN[k] = tickerResult.Data[0].Last
			//fmt.Printf("ticker = %v \n", tickerResult.Data)
			CoinPriceData := model.CoinPrice {
				Coin: v,
				Price: tickerResult.Data[0].Last,
			}
			cpChannel <- CoinPriceData
		} else {
			//fmt.Printf("ticker err = %v \n", tickerResult)
			coinErr := model.CoinError{
				Coin: v,
				Err: errors.New("errCode = " + strconv.Itoa(tickerResult.Code)),
			}
			errChannel <- coinErr
		}
		// 停10毫秒
		time.Sleep(time.Millisecond * 10)
	}
}

func Test () {
	fmt.Printf("testing \n")
	fmt.Printf("============================================== \n")
}