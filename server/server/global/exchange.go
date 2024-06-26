package global

import (
	"binanceNewCoin/server/model"
	"github.com/tidwall/gjson"
)

// 公共配置
var INIT_EXCHANGE_FINISH = false
var MAIN_QUEUE_TASK_RUNNING_STATUS = false

//  ======================= gate.io 配置 开始 =======================
const GATE_API_KEY = "xxxxxx"
const GATE_SECRET_KEY = "xxxxxxx"
const GATE_HOST = "https://api.gateio.ws"
var GATE_ALL_COIN = make([]string, 0)
var GATE_MONITOR_COIN = map[string]model.CoinItem{}
var LAST_UPDATE_GATE_MONITOR_TIME int64 // 上一次更新 GATE_MONITOR_COIN 的时间
var LAST_UPDATE_GATE_MONITOR_AMOUNT int // 上一次更新 GATE_MONITOR_COIN 成功的数量
//  ======================= gate.io 配置 结束 =======================

//  ======================= 币安 配置 开始 =======================
const BINANCE_API_KEY = "xxxxx"
const BINANCE_SECRET_KEY = "xxxxx"
const BIANACE_HOST = "https://api.binance.com"
var BINANCE_ALL_COINS = make([]gjson.Result, 0)
var BINANCE_COINS = make([]string, 0)
//  ======================= 币安 配置 结束 =======================

//  ======================= 抹茶 配置 开始 =======================
const MX_HOST = "https://www.mxc.com"
var MX_ALL_COIN = make([]string, 0)
//var MX_MONITOR_COIN = make([]gjson.Result, 0)
var MX_MONITOR_COIN = map[string]model.CoinItem{}
var MX_MONITOR_COIN_ARRAY = make([]string, 0)
var LAST_UPDATE_MX_MONITOR_TIME int64 // 上一次更新 MX_MONITOR_COIN 的时间
var LAST_UPDATE_MX_MONITOR_AMOUNT int // 上一次更新 MX_MONITOR_COIN 成功的数量
var MX_MONITOR_UPDATE_TIMES = 0		// 成功更新数
var MX_MONITOR_UPDATING = false		// 是否正在更新
var FIND_NEW_TITLE_TIME int64		// 发现新标题时间戳（毫秒）
var FINISH_ORDER_TIME int64			// 完成购买后的时间戳（毫秒）
//  ======================= 抹茶 配置 结束 =======================

