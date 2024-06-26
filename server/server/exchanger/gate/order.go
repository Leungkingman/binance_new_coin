package gate

import (
	"binanceNewCoin/server/global"
	"binanceNewCoin/server/logger"
	"binanceNewCoin/server/model"
	"binanceNewCoin/server/tool"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func GetOrderData(coin string, side string, amount string, price string, account string, is_auto_borrow bool) map[string]interface{} {
	if is_auto_borrow {
		orderData := map[string]interface{}{
			"currency_pair": coin + "_USDT",
			"side":          side,
			"amount":        amount,
			"account":       account,
			"price":         price,
			"auto_borrow":   true,
		}
		return orderData
	} else {
		orderData := map[string]interface{}{
			"currency_pair": coin + "_USDT",
			"side":          side,
			"amount":        amount,
			"account":       account,
			"price":         price,
		}
		return orderData
	}
}

func DoTrade(orderData map[string]interface{}) model.GateOrderResult {
	var gateOrderRes model.GateOrderResult
	const path = "/api/v4/spot/orders"
	const method = "POST"
	orderDataByte, _ := json.Marshal(orderData)
	orderDataJsonStr := string(orderDataByte)
	//queryString := tool.BuildOrderParams(orderData)
	//fmt.Printf("orderDataJsonStr = %v \n", orderDataJsonStr)
	hashedPayload := CreateHashPayload(orderDataJsonStr)
	//fmt.Printf("hashedPayload = %v \n", hashedPayload)
	signature := CreateHexQueryString(method, path, "", hashedPayload)
	//fmt.Printf("signature = %v \n", signature)
	_, jsonBody, _ := tool.ParseRequestParams(orderData)
	completeUrl := global.GATE_HOST + path
	req, _ := http.NewRequest(method, completeUrl, jsonBody)
	req.Header.Set("KEY", global.GATE_API_KEY)
	req.Header.Set("Timestamp", strconv.FormatInt(time.Now().Unix(), 10))
	req.Header.Set("SIGN", signature)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, httpErr := client.Do(req)
	if httpErr != nil {
		fmt.Printf("httpErr = %v \n", httpErr)
		logger.Addlog("请求错误："+httpErr.Error(), "error", "operate")
	}
	defer resp.Body.Close()
	httpResult, readHttpErr := ioutil.ReadAll(resp.Body)
	if readHttpErr != nil {
		fmt.Printf("readHttpErr = %v \n", readHttpErr)
	}
	_ = json.Unmarshal(httpResult, &gateOrderRes)
	logger.Addlog("请求结果：" + string(httpResult), "normal", "operate")
	//fmt.Printf("httpResult = %v \n", string(httpResult))
	return gateOrderRes
}

func MakeGateOrderStruct(or model.GateOrderResult, coinItem model.CoinItem) (dgo *model.DB_Gate_Order, err error) {
	total, err := tool.StringToFloat(or.Filled_total, 64)
	if err != nil {
		logger.Addlog("转换 Filled_total 失败"+err.Error(), "normal", "operate")
		return dgo, err
	}
	amount, err := tool.StringToFloat(or.Amount, 64)
	if err != nil {
		logger.Addlog("转换 Amount 失败"+err.Error(), "normal", "operate")
		return dgo, err
	}
	Prec := coinItem.Prec
	findTime := tool.Int64ToString(coinItem.Time)
	findPrice := strconv.FormatFloat(coinItem.Price, 'f', coinItem.Prec, 64)
	price, _ := decimal.NewFromFloat(total / amount).Round(int32(Prec + 1)).Float64() // 均价
	priceStr := strconv.FormatFloat(price, 'f', Prec, 64)
	orderData := &model.DB_Gate_Order{
		Order_id:    or.Id,
		Coin:        or.Fee_currency,
		Price:       priceStr,
		Find_price:  findPrice,
		Find_time:   findTime,
		Amount:      or.Amount,
		Total:       or.Filled_total,
		Left:        or.Left,
		Create_time: or.Create_time,
	}
	return orderData, nil
}

func MakeUpdateOrderStruct(or model.GateOrderResult) (dgo *model.DB_Gate_Order, err error) {

	coinItem, ok := global.GATE_MONITOR_COIN[or.Fee_currency]
	if ok {
		total, _ := tool.StringToFloat(or.Filled_total, 64)
		amount, _ := tool.StringToFloat(or.Amount, 64)
		Prec := coinItem.Prec
		findTime := or.Create_time + "000"
		findPrice := "0"
		price, _ := decimal.NewFromFloat(total / amount).Round(int32(Prec + 1)).Float64()
		priceStr := strconv.FormatFloat(price, 'f', Prec, 64)
		orderData := &model.DB_Gate_Order{
			Order_id:    or.Id,
			Coin:        or.Fee_currency,
			Price:       priceStr,
			Find_price:  findPrice,
			Find_time:   findTime,
			Amount:      or.Amount,
			Total:       or.Filled_total,
			Left:        or.Left,
			Create_time: or.Create_time,
		}
		return orderData, nil
	} else {
		return dgo, errors.New("gate.io 不存在该币种")
	}
}

func CreateGateOrder(o model.DB_Gate_Order) (err error, order model.DB_Gate_Order) {
	err = global.DB_ENGINE.Create(&o).Error
	return err, o
}

func CreateHashPayload(orderDataJsonStr string) string {
	h := sha512.New()
	h.Write([]byte(orderDataJsonStr))
	return hex.EncodeToString(h.Sum(nil))
}

func CreateHexQueryString(method, path, queryString, hashedPayload string) string {
	timeStamp := strconv.FormatInt(time.Now().Unix(), 10)
	str := method + "\n" + path + "\n" + queryString + "\n" + hashedPayload + "\n" + timeStamp
	h := hmac.New(sha512.New, []byte(global.GATE_SECRET_KEY))
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func FetchGateOrder(orderId, pair string) model.GateOrderResult {
	var gateOrderRes model.GateOrderResult
	path := "/api/v4/spot/orders/" + orderId
	const method = "GET"
	query_param := "currency_pair=" + pair + "_USDT"
	hashedPayload := CreateHashPayload("")
	//fmt.Printf("hashedPayload = %v \n ", hashedPayload)
	signature := CreateHexQueryString(method, path, query_param, hashedPayload)
	//fmt.Printf("signature = %v \n", signature)
	path = path + "?" + query_param
	completeUrl := global.GATE_HOST + path
	req, _ := http.NewRequest(method, completeUrl, nil)
	req.Header.Set("KEY", global.GATE_API_KEY)
	req.Header.Set("Timestamp", strconv.FormatInt(time.Now().Unix(), 10))
	req.Header.Set("SIGN", signature)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	//fmt.Printf("req = %v \n", req)

	client := &http.Client{}
	resp, httpErr := client.Do(req)
	if httpErr != nil {
		//fmt.Printf("httpErr = %v \n", httpErr)
		logger.Addlog("请求错误："+httpErr.Error(), "error", "operate")
	}
	defer resp.Body.Close()
	httpResult, readHttpErr := ioutil.ReadAll(resp.Body)
	if readHttpErr != nil {
		fmt.Printf("readHttpErr = %v \n", readHttpErr)
	}
	_ = json.Unmarshal(httpResult, &gateOrderRes)
	logger.Addlog("请求结果："+string(httpResult), "normal", "operate")
	//fmt.Printf("httpResult = %v \n", string(httpResult))
	return gateOrderRes
}

func TransfersAccount(transferBody model.TransfersBody) interface{} {
	const path = "/api/v4/wallet/transfers"
	const method = "POST"
	transferByte, _ := json.Marshal(transferBody)
	transferDataJsonStr := string(transferByte)
	hashedPayload := CreateHashPayload(transferDataJsonStr)
	signature := CreateHexQueryString(method, path, "", hashedPayload)
	//fmt.Printf("signature = %v \n", signature)
	_, jsonBody, _ := tool.ParseRequestParams(transferBody)
	completeUrl := global.GATE_HOST + path
	req, _ := http.NewRequest(method, completeUrl, jsonBody)
	req.Header.Set("KEY", global.GATE_API_KEY)
	req.Header.Set("Timestamp", strconv.FormatInt(time.Now().Unix(), 10))
	req.Header.Set("SIGN", signature)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, httpErr := client.Do(req)
	if httpErr != nil {
		//fmt.Printf("httpErr = %v \n", httpErr)
		logger.Addlog("请求错误："+httpErr.Error(), "error", "operate")
	}
	defer resp.Body.Close()
	httpResult, readHttpErr := ioutil.ReadAll(resp.Body)
	if readHttpErr != nil {
		fmt.Printf("readHttpErr = %v \n", readHttpErr)
	}
	fmt.Printf("TransfersAccount httpResult = %v \n", string(httpResult))
	return httpResult
}
