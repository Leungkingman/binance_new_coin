package api

import (
	"binanceNewCoin/server/exchanger/binance"
	"binanceNewCoin/server/exchanger/gate"
	"binanceNewCoin/server/global"
	"binanceNewCoin/server/global/response"
	"github.com/gin-gonic/gin"
)

func InitialExchangeData(c *gin.Context) {
	if global.INIT_EXCHANGE_FINISH {
		response.OkWithMessage("初始化进行中或已完成", c)
		return
	}
	response.OkWithMessage("初始化进行中", c)
	// ==================== 测试代码 开始 ====================
	//global.MX_MONITOR_COIN_ARRAY = append(global.MX_MONITOR_COIN_ARRAY, "KLAY")
	//global.MX_MONITOR_COIN["KLAY"] = model.CoinItem {
	//	Price: 0,
	//	Prec: 0,
	//}
	// ==================== 测试代码 结束 ====================
	go func() {
		gate.SetGateCoinPrice()
	}()
}

func GetExchangeMonitorCoin(c *gin.Context) {
	var exchangeCoin map[string]interface{}

	// mx 交易所数据
	var mxData map[string]interface{}
	mxData["update_amount"] = global.LAST_UPDATE_MX_MONITOR_AMOUNT
	mxData["last_update_time"] = global.LAST_UPDATE_MX_MONITOR_TIME
	mxData["coinData"] = global.MX_MONITOR_COIN

	var gateData map[string]interface{}
	gateData["update_amount"] = global.LAST_UPDATE_GATE_MONITOR_TIME
	gateData["last_update_time"] = global.LAST_UPDATE_GATE_MONITOR_AMOUNT
	gateData["coinData"] = global.GATE_MONITOR_COIN

	exchangeCoin["mx"] = mxData
	exchangeCoin["gate"] = gateData
	exchangeCoin["newTitle"] = binance.GetCurrentNewTitle()

	response.Result(0, exchangeCoin,"", c)
}

func GetBinanceNewTitle(c *gin.Context) {
	data := map[string]interface{}{
		"title": binance.GetCurrentNewTitle(),
	}
	response.Result(0, data,"", c)
}

func GetInitialInfo(c *gin.Context) {
	result := map[string]interface{}{
		"exchange_log":  global.EXCHANGE_LOGS,
	}
	response.Result(0, result, "", c)
}