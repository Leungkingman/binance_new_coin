package mx

import (
	"binanceNewCoin/server/global"
	"binanceNewCoin/server/model"
	"binanceNewCoin/server/tool"
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Mx struct {
	AccessKey    string
	SecretKey    string
	UsdtUse      float64
	SecondBuy    float64
	Slippage     float64
	NextSlippage float64
	BuyTime      int
}

type MarketItem struct {
	Symbol      string `json:"symbol"`
	Volume      string `json:"volume"`
	High        string `json:"high"`
	Low         string `json:"low"`
	Bid         string `json:"bid"`
	Ask         string `json:"ask"`
	Open        string `json:"open"`
	Last        string `json:"last"`
	Time        int    `json:"time"`
	Change_rate string `json:"change_rate"`
}

type GetResult struct {
	Code int          `json:"code"`
	Data []MarketItem `json:"data"`
}

type OrderResult struct {
	Code int    `json:"code"`
	Data string `json:"data"`
}

func (mx *Mx) GetTickerBuyData(coin string) (string, string, error) {
	lastPrice := ""
	lastAmount := ""
	url := global.MX_HOST + "/open/api/v2/market/ticker?symbol=" + coin + "_USDT"
	//fmt.Printf("url = %v \n", url)
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return lastPrice, lastAmount, err
	}
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return lastPrice, lastAmount, err
	}

	var tickerResult GetResult
	err = json.Unmarshal(result, &tickerResult)
	if err != nil {
		return lastPrice, lastAmount, err
	}

	//fmt.Printf("code = %v \n", tickerResult.Code)
	//fmt.Printf("data = %v \n", tickerResult.Data)

	if tickerResult.Code == 200 && len(tickerResult.Data) > 0 {
		tempPriceFloat, _ := strconv.ParseFloat(tickerResult.Data[0].Last, 64)
		tempPrice, _ := decimal.NewFromFloat(tempPriceFloat).Round(3).Float64()
		var orderPrice float64
		var amount float64
		//fmt.Printf("BuyTime = %v \n", mx.BuyTime)
		if mx.BuyTime == 1 {
			orderPrice = tempPrice * (1 + mx.Slippage/100)
			amount, _ = decimal.NewFromFloat(mx.UsdtUse / orderPrice).Round(3).Float64()
			//fmt.Printf("orderPrice 111 = %v \n", orderPrice)
			//fmt.Printf("tempPrice 111 = %v \n", tempPrice)
			//fmt.Printf("Slippage 111 = %v \n", mx.Slippage)
			//fmt.Printf("UsdtUse 111 = %v \n", mx.UsdtUse)
			//fmt.Printf("amount 111 = %v \n", amount)
		} else if mx.BuyTime == 2 {
			orderPrice = tempPrice * (1 + mx.NextSlippage/100)
			amount, _ = decimal.NewFromFloat(mx.SecondBuy / orderPrice).Round(3).Float64()
			//fmt.Printf("orderPrice 222 = %v \n", orderPrice)
			//fmt.Printf("tempPrice 222 = %v \n", tempPrice)
			//fmt.Printf("NextSlippage 222 = %v \n", mx.NextSlippage)
			//fmt.Printf("SecondBuy 222 = %v \n", mx.SecondBuy)
			//fmt.Printf("amount 222 = %v \n", amount)
		}
		lastPrice = strconv.FormatFloat(orderPrice, 'f', 2, 64)
		lastAmount = strconv.FormatFloat(amount, 'f', 2, 64)
	}

	//fmt.Printf("lastPrice = %v \n", lastPrice)
	//fmt.Printf("lastAmount = %v \n", lastAmount)
	return lastPrice, lastAmount, nil
}

func (mx *Mx) OrderBuy(orderBody map[string]string) (map[string]interface{}, error) {
	const path = "/open/api/v2/order/place"
	const method = "POST"
	result := map[string]interface{}{}

	timeStamp := strconv.FormatInt(time.Now().Unix(), 10)
	params := map[string]string{
		"api_key":     mx.AccessKey,
		"req_time":    timeStamp,
		"recv_window": "60",
	}

	queryStringBeforeSign := tool.BuildOrderParams(params)
	signature := tool.CreateSignature(method, path, queryStringBeforeSign, mx.SecretKey)
	params["sign"] = signature
	queryString := tool.BuildOrderParams(params)
	_, jsonBody, _ := tool.ParseRequestParams(orderBody)

	completeUrl := global.MX_HOST + path + "?" + queryString
	req, _ := http.NewRequest(method, completeUrl, jsonBody)
	req.Header.Set("Content-Type", "application/json")
	//printRequest(req, jsonStr)

	client := &http.Client{}
	resp, httpErr := client.Do(req)
	if httpErr != nil {
		return result, httpErr
	}
	defer resp.Body.Close()
	httpResult, readHttpErr := ioutil.ReadAll(resp.Body)
	if readHttpErr != nil {
		return result, readHttpErr
	}
	var orderRes OrderResult
	_ = json.Unmarshal(httpResult, &orderRes)
	//fmt.Printf("orderRes code = ", orderRes.Code)
	//fmt.Printf("\n")
	//fmt.Printf("orderRes data = ", orderRes.Data)
	//fmt.Printf("\n")
	//fmt.Printf("======================================= \n")
	result["code"] = orderRes.Code
	result["data"] = orderRes.Data
	return result, nil
}

func (mx *Mx) GetOrderData(order_ids string) (model.MxOrderResult, error) {
	const path = "/open/api/v2/order/query"
	const method = "GET"
	var mxOrderRes model.MxOrderResult

	timeStamp := strconv.FormatInt(time.Now().Unix(), 10)
	params := map[string]string{
		"api_key":     mx.AccessKey,
		"req_time":    timeStamp,
		"recv_window": "60",
		"order_ids":   order_ids,
	}

	queryStringBeforeSign := tool.BuildOrderParams(params)
	signature := tool.CreateSignature(method, path, queryStringBeforeSign, mx.SecretKey)
	params["sign"] = signature
	queryString := tool.BuildOrderParams(params)

	completeUrl := global.MX_HOST + path + "?" + queryString
	//fmt.Printf("completeUrl = ", completeUrl)
	req, _ := http.NewRequest(method, completeUrl, nil)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, httpErr := client.Do(req)
	if httpErr != nil {
		return mxOrderRes, httpErr
	}
	defer resp.Body.Close()
	httpResult, readHttpErr := ioutil.ReadAll(resp.Body)
	if readHttpErr != nil {
		return mxOrderRes, readHttpErr
	}
	_ = json.Unmarshal(httpResult, &mxOrderRes)
	//fmt.Printf("mxOrderRes code = ", mxOrderRes.Code)
	//fmt.Printf("\n")
	//fmt.Printf("mxOrderRes data = ", mxOrderRes.Data)
	//fmt.Printf("\n")
	//fmt.Printf("======================================= \n")
	return mxOrderRes, nil
}

func printRequest(request *http.Request, body string) {
	fmt.Printf(" ================== Request start ================== ")
	fmt.Println("Url: " + request.URL.String())
	fmt.Println("Method: " + strings.ToUpper(request.Method))
	if len(request.Header) > 0 {
		fmt.Println("\tHeaders: ")
		for k, v := range request.Header {
			fmt.Println("\t\t" + k + ": " + v[0])
		}
	}
	fmt.Println("Body: " + body)
	fmt.Printf(" ================== Request end ================== ")
}
