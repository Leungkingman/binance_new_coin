package gate

import (
	"binanceNewCoin/server/global"
	"binanceNewCoin/server/logger"
	"binanceNewCoin/server/model"
	"binanceNewCoin/server/tool"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func GetAllCoin () ([]string, error) {
	var coins = make([]string, 0)
	completeUrl := global.GATE_HOST + "/api/v4/spot/currencies"
	resp, err := http.Get(completeUrl)
	defer resp.Body.Close()
	if err != nil {
		fmt.Printf("err 00 = ", err.Error())
		return coins, err
	} else {
		result, _ := ioutil.ReadAll(resp.Body)
		var gateCurrencyData []model.GateCurrenciesItem
		err = json.Unmarshal(result, &gateCurrencyData)
		fmt.Printf("result = %v \n", string(result))
		if err != nil {
			fmt.Printf("err 11 = %v \n", err.Error())
			return coins, err
		}
		//fmt.Printf("result = %v \n", result)
		for _, item := range gateCurrencyData {
			if !item.Delisted && !strings.Contains(item.Currency, "_") {
				coins = append(coins, item.Currency)
			}
		}
	}
	return coins, nil
}

func SetGateCoinPrice() {
	logger.ClearNormmalLog("exchange")
	success := 0
	completeUrl := global.GATE_HOST + "/api/v4/spot/tickers"
	resp, err := http.Get(completeUrl)
	if err != nil {
		fmt.Printf("http err = ", err.Error())
	} else {
		result, _ := ioutil.ReadAll(resp.Body)
		//fmt.Printf("tickers result = %v \n", string(result))
		var gateTickerData []model.GateTickerItem
		err = json.Unmarshal(result, &gateTickerData)
		if err != nil {
			fmt.Printf("json err = %v \n", err.Error())
		}
		//fmt.Printf("result = %v \n", result)
		//fmt.Printf("result = %v \n", string(result))
		for _, item := range gateTickerData {
			if item.Last != "0" {
				//fmt.Printf("value = %v \n", global.GATE_MONITOR_COIN[coin].Price)
				success += 1
				priceStr := item.Last
				price, _ := strconv.ParseFloat(priceStr,64)
				prec := 0
				startIndex := strings.Index(priceStr, ".")
				priceLength := len(priceStr)
				prec = len(priceStr[startIndex + 1: priceLength])
				global.GATE_MONITOR_COIN[item.Currency_pair] = model.CoinItem{
					Price: price,
					Prec: prec,
				}
				//fmt.Printf("CoinItem = %v \n", global.GATE_MONITOR_COIN[coin])
			}
		}
		global.LAST_UPDATE_GATE_MONITOR_TIME = time.Now().Unix()
		global.LAST_UPDATE_GATE_MONITOR_AMOUNT = success
		//fmt.Printf("====================================================================== \n")
		//fmt.Printf("GATE_MONITOR_COIN = %v \n", global.GATE_MONITOR_COIN)

		//number := 1
		//for k, v := range global.GATE_MONITOR_COIN {
		//	price := strconv.FormatFloat(v.Price, 'f', v.Prec, 64)
		//	logMsg := tool.IntToString(number) + ". " + k + "：" + price
		//	logger.Addlog(logMsg, "normal", "exchange")
		//	number += 1
		//}
		logger.Addlog("币种个数：" + tool.IntToString(len(global.GATE_MONITOR_COIN)), "normal", "exchange")
		global.INIT_EXCHANGE_FINISH = true
		defer resp.Body.Close()
	}
}

func GetGateMarginCoins () {
	completeUrl := global.GATE_HOST + "/margin/currency_pairs"
	resp, err := http.Get(completeUrl)
	if err != nil {
		fmt.Printf("http err = ", err.Error())
	} else {
		result, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("tickers result = %v \n", string(result))
		defer resp.Body.Close()
	}
}
