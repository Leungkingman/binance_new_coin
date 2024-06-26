package binance

import (
	"binanceNewCoin/server/global"
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func GetAll () ([]gjson.Result, error) {
	var resultData []gjson.Result
	params := map[string]string{
		"timestamp": strconv.FormatInt(time.Now().UnixNano() / 1e6,10),
	}
	queryString := GetCompleteQueryString(params)
	completeUrl := global.BIANACE_HOST + "/sapi/v1/capital/config/getall?" + queryString
	//fmt.Printf("completeUrl = %v \n", completeUrl)
	req, err := http.NewRequest("GET", completeUrl, nil)
	if err != nil {
		return resultData, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-MBX-APIKEY", global.BINANCE_API_KEY)
	client := &http.Client{}
	resp, httpErr := client.Do(req)
	if httpErr != nil {
		//fmt.Printf("httpErr = %v \n", httpErr)
		return resultData, err
	}
	defer resp.Body.Close()
	result, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		//fmt.Printf("httpErr = %v \n", httpErr)
		return resultData, readErr
	}
	return gjson.Get(string(result), "#.coin").Array(), nil
	//fmt.Printf("BINANCE_ALL_COINS = %v \n", global.BINANCE_ALL_COINS)
}

func GetAllCoin () ([]string, error) {
	//fmt.Printf("start time = %v \n", time.Now().UnixNano() / 1e6)
	resultData := make([]string, 0)
	params := map[string]string{
		"timestamp": strconv.FormatInt(time.Now().UnixNano() / 1e6,10),
	}
	queryString := GetCompleteQueryString(params)
	completeUrl := global.BIANACE_HOST + "/sapi/v1/capital/config/getall?" + queryString
	req, err := http.NewRequest("GET", completeUrl, nil)
	if err != nil {
		return resultData, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-MBX-APIKEY", global.BINANCE_API_KEY)
	client := &http.Client{}
	resp, httpErr := client.Do(req)
	if httpErr != nil {
		fmt.Printf("httpErr 111 = %v \n", httpErr)
		return resultData, httpErr
	}
	//fmt.Printf("http end time = %v \n", time.Now().UnixNano() / 1e6)
	defer resp.Body.Close()
	result, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		fmt.Printf("readErr = %v \n", readErr)
		return resultData, readErr
	}
	//fmt.Printf("BINANCE_ALL_COINS = %v \n", global.BINANCE_ALL_COINS)
	//return string(result), nil
	gjsonResult := gjson.Get(string(result), "#.coin").Array()
	for _, v := range gjsonResult {
		resultData = append(resultData, v.Str)
	}
	//fmt.Printf("end time = %v \n", time.Now().UnixNano() / 1e6)
	return resultData, nil
}

func CoinInBinanceCoin(coin string, array []gjson.Result) bool {
	coinExist := false
	for _, v := range array {
		if v.Str == coin {
			coinExist = true
			break
		}
	}
	return coinExist
}
