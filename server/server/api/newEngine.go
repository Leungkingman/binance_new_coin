package api

import (
	"binanceNewCoin/server/exchanger/binance"
	"binanceNewCoin/server/exchanger/gate"
	"binanceNewCoin/server/global"
	"binanceNewCoin/server/global/response"
	"binanceNewCoin/server/informer"
	"binanceNewCoin/server/logger"
	"binanceNewCoin/server/middleware"
	"binanceNewCoin/server/model"
	"binanceNewCoin/server/model/request"
	"binanceNewCoin/server/tool"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var dbDomains []model.DB_Domain
var dbArticles []model.Article
var dbCoin model.Coin
var ips []string
var server_queue_time_gap int
var buyUsdt string
var slippage string

// 主服务器轮询任务
// 逻辑：
// 1. 从参数从提取服务器ips
// 2. 启动轮询任务，轮询任务逻辑：
// 2.1	从数据库中读取轮序间隔时间 t
// 2.2  for 循环请求各个服务器的接口，然后间隔 t 时间
// 2.3  如果轮询完一组，就从头开始

func StartMainQueueTask(c *gin.Context) {
	if global.MAIN_QUEUE_TASK_RUNNING_STATUS {
		response.Result(0, nil, "服务器已启动", c)
		return
	}
	if !global.INIT_EXCHANGE_FINISH {
		response.Result(-1, nil, "请先初始化系统", c)
		return
	}
	var MainQueueTaskData request.MainQueueTaskData
	_ = c.ShouldBindJSON(&MainQueueTaskData)
	ips = MainQueueTaskData.Ips

	claims := c.MustGet("claims").(*middleware.CustomClaims)
	uid := claims.Uid
	global.DB_ENGINE.Table("bm_user").First(&user, uid)
	// 获取币种
	global.DB_ENGINE.Table("bm_exist_coin").First(&coin)
	coins := strings.Split(coin.Coin, ",")
	server_queue_time_gap, _ = strconv.Atoi(user.Server_queue_time_gap)
	buyUsdt = user.Buy_usdt
	slippage = user.Slippage
	// 获取标题s
	existTitles := make([]string, 0)
	global.DB_ENGINE.Table("bm_articles").Find(&articles)
	for i := 0; i < len(articles); i++ {
		existTitles = append(existTitles, articles[i].Content)
	}
	// 获取域名
	existDomains := make([]string, 0)
	global.DB_ENGINE.Table("bm_domain").Find(&dbDomains)
	for i := 0; i < len(dbDomains); i++ {
		existDomains = append(existDomains, dbDomains[i].Domain)
	}

	global.MAIN_QUEUE_TASK_RUNNING_STATUS = true
	go func() {
		QueueTask(ips, server_queue_time_gap, existDomains, existTitles, coins, buyUsdt, slippage)
	}()
	response.Result(0, nil, "服务器启动", c)
}

func QueueTask(ips []string, stopTime int, domains []string, existTitles []string, coins []string, buyUsdt string, slippage string) {
	index := 0
	for {
		if (global.MAIN_QUEUE_TASK_RUNNING_STATUS) {
			ip := ips[index]
			url := "http://" + ip + ":8880/engine/fetchBinanceData"
			postDataValue := map[string]interface{} {
				"domains": domains,
				"titles": existTitles,
				"coins": coins,
				"buyUsdt": buyUsdt,
				"slippage": slippage,
			}
			postData, _ := json.Marshal(postDataValue)
			resp, err := http.Post(url, "application/json", bytes.NewBuffer(postData))
			if err != nil {
				logger.Addlog("QueueTask err = " + err.Error(),"normal", "error")
				logger.Addlog("ip: " + ip, "normal", "error")
				continue
			}
			defer resp.Body.Close()
			if index < len(ips)-1 {
				index += 1
			} else {
				index = 0
			}
			time.Sleep(time.Duration(stopTime) * time.Millisecond)
		} else {
			break
		}
	}
}

func StopQueueTaskRunningStatus(c *gin.Context) {
	global.MAIN_QUEUE_TASK_RUNNING_STATUS = false
	response.Result(0, nil, "更新成功", c)
}

func StopQueueTask(c *gin.Context) {
	url := "http://8.218.81.24:8880/engine/stopQueueTaskRunningStatus"
	resp, err := http.Get(url)
	if err != nil {
		logger.Addlog("StopQueueTask err = " + err.Error(),"normal", "error")
	}
	defer resp.Body.Close()
	response.Result(0, nil, "主服务器任务已暂停", c)
}

func RestartMainQueueByServer(c *gin.Context) {
	url := "http://8.218.81.24:8880/engine/startMainQueueWithoutPermission"
	resp, err := http.Get(url)
	if err != nil {
		logger.Addlog("StopQueueTask err = " + err.Error(),"normal", "error")
	}
	defer resp.Body.Close()
	response.Result(0, nil, "主服务器任务已暂停", c)
}

func StartMainQueueWithoutPermission(c *gin.Context) {
	time.Sleep(time.Duration(server_queue_time_gap) * time.Millisecond + 1000)
	// 获取币种
	global.DB_ENGINE.Table("bm_exist_coin").First(&coin)
	coins := strings.Split(coin.Coin, ",")
	// 获取标题s
	existTitles := make([]string, 0)
	global.DB_ENGINE.Table("bm_articles").Find(&articles)
	for i := 0; i < len(articles); i++ {
		existTitles = append(existTitles, articles[i].Content)
	}
	// 获取域名
	existDomains := make([]string, 0)
	global.DB_ENGINE.Table("bm_domain").Find(&dbDomains)
	for i := 0; i < len(dbDomains); i++ {
		existDomains = append(existDomains, dbDomains[i].Domain)
	}

	global.MAIN_QUEUE_TASK_RUNNING_STATUS = true
	go func() {
		QueueTask(ips, server_queue_time_gap, existDomains, existTitles, coins, buyUsdt, slippage)
	}()
	response.Result(0, nil, "服务器启动", c)
}

func GetQueueTaskRunningStatus(c *gin.Context) {
	response.Result(0, global.MAIN_QUEUE_TASK_RUNNING_STATUS, "", c)
}

// 从服务器请求任务
// 清除执行日志及定时任务日志
func FetchBinanceData(c *gin.Context) {
	var postData request.SingleServerFetchTask
	_ = c.ShouldBindJSON(&postData)
	existDomains := postData.Domains
	binanceTitles := postData.Titles
	ExistCoins := postData.Coins
	p_buyUsdt, usdtErr := tool.StringToFloat(postData.BuyUsdt, 64)
	if usdtErr != nil {
		response.FailWithMessage("usdtErr: " + usdtErr.Error(), c)
		return
	}
	p_slippage, slippageErr := tool.StringToFloat(postData.Slippage, 64)
	if slippageErr != nil {
		response.FailWithMessage("slippageErr: " + slippageErr.Error(), c)
		return
	}

	logger.ClearNormmalLog("mobileApi")
	logger.ClearNormmalLog("exchange")

	// 获取 gate.io 交易对数据
	go func() {
		gate.SetGateCoinPrice()
	}()

	go func() {
		for i := 0; i < len(existDomains); i++ {
			if i == 0 {
				logger.Addlog("开始轮询时间：" + tool.Int64ToString(time.Now().Unix()), "normal", "mobileApi")
			}
			domain := existDomains[i]
			success, _, title, createTime, _ := binance.FetchFromApi(0, domain)
			if success {
				logger.Addlog("timeStamp："+tool.Int64ToString(createTime), "normal", "mobileApi")
				logger.Addlog(title, "normal", "mobileApi")
				isEixst := tool.StringInArray(title, binanceTitles)
				if !isEixst {
					logger.Addlog("发现新标题：" + tool.Int64ToString(time.Now().UnixNano() / 1e6), "normal", "operate")
					logger.Addlog("标题发布时间：" + tool.Int64ToString(createTime), "normal", "operate")
					logger.Addlog("domain："+domain, "normal", "operate")
					logger.Addlog("新标题："+title, "normal", "operate")
					coins := HandleNewTitle(title)                                     // 从标题中提取 coin，可能存在多个
					find, newCoin, usdtCoinItem, _, checkMsg := HandleGateCoins(coins, ExistCoins) // 从 gate.io 的币中查找这个币是否存在
					logger.Addlog(checkMsg, "normal", "operate")
					if find {
						logger.Addlog("newCoin = "+newCoin, "normal", "operate")
						tradeResult := DoGateTrade(newCoin, usdtCoinItem, p_buyUsdt, p_slippage)
						HandleTradeResult(tradeResult, usdtCoinItem)

						logger.Addlog("更新币数据表", "normal", "operate")
						// ====================== 添加新币到数据库 start ======================
						global.DB_ENGINE.Table("bm_exist_coin").First(&dbCoin)
						if dbCoin.Coin == "" {
							dbCoin.Coin = newCoin
						} else {
							completeCoinData := dbCoin.Coin + "," + newCoin
							dbCoin.Coin = completeCoinData
						}
						global.DB_ENGINE.Table("bm_exist_coin").Save(&dbCoin)
						// ====================== 添加新币到数据库 end ======================
						UpdateBinanceTitle(title, 20)
						// 停止主服务器轮询
						StopQueueTask(c)
					}
					UpdateBinanceTitle(title, 10)
					// 停止主服务器轮询，并重新开始
					RestartMainQueueByServer(c)
				}
			} else {
				logger.ClearNormmalLog("error")
				logger.Addlog("domain："+domain, "error", "error")
				logger.Addlog(title, "error", "error")
			}
		}
	}()
	response.Result(0, nil, "执行完毕", c)
}

func HandleNewTitle(title string) (coins []string) {
	var matchCoins []string
	if strings.Contains(title, "Binance Will List") {
		matchCoins = coinReg.FindStringSubmatch(title)
	}
	return matchCoins
}

func HandleGateCoins(coins []string, existCoins []string) (bool, string, model.CoinItem, model.CoinItem, string) {
	find := false
	targetCoin := ""
	msg := ""
	var usdtCoinResult model.CoinItem
	var ethCoinResult model.CoinItem
	for _, v := range coins {
		targetCoin = v
		usdtValue, ok := global.GATE_MONITOR_COIN[v+"_USDT"]
		if ok {
			isEixst := tool.StringInArray(targetCoin, existCoins)
			if !isEixst {
				msg = "找到目标币种：" + targetCoin
				find = true
				usdtCoinResult = usdtValue
				break
			} else {
				msg = "币种已存在：" + targetCoin + "；不进行后续操作"
			}
		} else {
			msg = "平台不存在该币种：" + targetCoin + "；不进行后续操作"
		}
	}
	return find, targetCoin, usdtCoinResult, ethCoinResult, msg
}

func DoGateTrade(coin string, coinPriceData model.CoinItem, buyUsdt float64, slippage float64) model.GateOrderResult {
	coinPrice := strconv.FormatFloat(coinPriceData.Price, 'f', coinPriceData.Prec, 64)
	tempPrice := coinPriceData.Price * (1 + slippage/100)
	tempQuantity, _ := decimal.NewFromFloat(buyUsdt / tempPrice).Round(3).Float64()
	quantity := strconv.FormatFloat(tempQuantity, 'f', 2, 64) // 转换 string

	BuyUsdtStr := strconv.FormatFloat(buyUsdt, 'f', 3, 64)
	price := strconv.FormatFloat(tempPrice, 'f', 2, 64)
	logMsg := "购买数量：使用usdt = " + BuyUsdtStr + "；数量 = " + quantity + "；价格 = " + price + "；发现价：" + coinPrice
	logger.Addlog(logMsg, "normal", "operate")

	orderPrice := strconv.FormatFloat(tempPrice, 'f', coinPriceData.Prec, 64)
	gateOrderData := gate.GetOrderData(coin, "buy", quantity, orderPrice, "spot", false)
	//fmt.Printf("gateOrderData = %v \n", gateOrderData)
	// 接收交易结果
	tradeResult := gate.DoTrade(gateOrderData)
	tradeResultStr, _ := json.Marshal(tradeResult)
	logger.Addlog("请求结果："+string(tradeResultStr), "normal", "operate")
	return tradeResult
}

func HandleTradeResult(tradeRes model.GateOrderResult, coinItem model.CoinItem) {
	if tradeRes.Id != "" {
		if tradeRes.Left == "0" {
			logger.Addlog("Fit!! 全部成交", "normal", "operate")
		} else {
			logger.Addlog("部分成交或没有成交。剩余 "+tradeRes.Left+" 未成交。", "normal", "operate")
		}
		gos, makeErr := gate.MakeGateOrderStruct(tradeRes, coinItem)
		if makeErr == nil {
			err := global.DB_ENGINE.Table("bm_order").Create(gos).Error
			if err != nil {
				logger.Addlog("添加订单结果到数据库失败："+err.Error(), "normal", "operate")
			}
		}
		// 短信通知
		logger.Addlog("发送短信通知、更新标题，流程结束", "normal", "operate")
		informer.InformationNewListing(global.GATE_ORDER_TPL_CODE, 5)
	} else {
		logger.Addlog("没成交", "normal", "operate")
	}
}

func UpdateBinanceTitle(title string, operateType int) {
	isAddNew := binance.UpdateNewTitle(title)
	logger.Addlog("isAddNew = "+strconv.FormatBool(isAddNew), "normal", "operate")
	if isAddNew && operateType == 10 {
		logger.Addlog("新标题通知", "normal", "operate")
		informer.InformationNewListing(global.NEW_LISTING_INFO_TPL_CODE, 1)
	}
}
