package engine

import (
	"binanceNewCoin/server/exchanger/binance"
	"binanceNewCoin/server/exchanger/gate"
	"binanceNewCoin/server/exchanger/mx"
	"binanceNewCoin/server/global"
	"binanceNewCoin/server/informer"
	"binanceNewCoin/server/logger"
	"binanceNewCoin/server/model"
	"binanceNewCoin/server/tool"
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Concurrent struct {
	WorkerCount       int
	KeepRunning       bool
	ExistCoin         []string
	ExistTitles       []string
	Domains           []string
	BuyUsdt           float64
	Slippage          float64
	RequestTimeGap    int
	NewRequestTimeGap int
	HasOrder          bool
}

type NewConcurrent struct {
	ips []string
}

var Mex *mx.Mx
var db_coin model.Coin
var mxOrder model.MxOrder
var article model.Article

var coinReg = regexp.MustCompile(`\((.*?)\)`)
var currentNewTitle string

func (c *Concurrent) Run() {
	var apiLogChannel = make(chan model.LogData)
	currentNewTitle = binance.GetCurrentNewTitle()
	c.StartFetchTask(apiLogChannel)
	go func() { // 每分钟清一次 normal 日志
		for {
			logger.ClearNormmalLog("mobileApi")
			logger.ClearNormmalLog("error")
			time.Sleep(time.Minute * 1)
		}
	}()

	go func() {
		c.StartFetchGateData()
	}()

	for {
		logData, ok := <-apiLogChannel
		if !ok || !c.KeepRunning {
			logger.Addlog("运行暂停", "error", "error")
			break
		}
		//fmt.Printf("=============================================== \n")
		domain := c.Domains[logData.DomainIndex]
		if logData.Success {
			logger.Addlog(domain+"；", "normal", "mobileApi")
			logger.Addlog("timeStamp：" + tool.Int64ToString(logData.CreateTime), "normal", "mobileApi")
			title := logData.Content
			logger.Addlog(title, "normal", "mobileApi")
			isEixst := tool.StringInArray(title, c.ExistTitles)
			if !isEixst {
				logger.Addlog("domain："+domain, "normal", "operate")
				logger.Addlog("新标题："+title, "normal", "operate")
				coins := c.HandleNewTitle(title)                                  // 从标题中提取 coin，可能存在多个
				find, coin, usdtCoinItem, _, checkMsg := c.HandleGateCoins(coins) // 从 gate.io 的币中查找这个币是否存在
				logger.Addlog(checkMsg, "normal", "operate")
				logger.Addlog("标题发布时间：", "normal", "operate")
				if find && c.KeepRunning {
					c.Stop()
					logger.Addlog("coin = "+coin, "normal", "operate")
					tradeResult := c.DoGateTrade(coin, usdtCoinItem)
					//c.DoEthTrade(ethCoinItem)
					c.HandleTradeResult(tradeResult, usdtCoinItem)

					logger.Addlog("更新币数据表", "normal", "operate")
					// ====================== 添加新币到数据库 start ======================
					global.DB_ENGINE.Table("bm_exist_coin").First(&db_coin)
					if db_coin.Coin == "" {
						db_coin.Coin = coin
					} else {
						completeCoinData := db_coin.Coin + "," + coin
						db_coin.Coin = completeCoinData
					}
					global.DB_ENGINE.Table("bm_exist_coin").Save(&db_coin)
					// ====================== 添加新币到数据库 end ======================
					c.UpdateBinanceTitle(title, 20)
					break
				}
				c.UpdateBinanceTitle(title, 10)
			}
		} else {
			logger.Addlog("domain："+domain, "error", "error")
			logger.Addlog(logData.Content, "error", "error")
			//if c.KeepRunning {
			//	c.Stop()
			//	close(apiLogChannel)
			//	time.Sleep(time.Minute * 40)
			//	logger.Addlog("程序重启", "error", "error")
			//	go func() {
			//		c.Run()
			//	}()
			//} else {
			//	logger.Addlog("程序结束", "error", "error")
			//}
			//break
		}
	}
}

//func (c *Concurrent) RunNewTask() {
//	startTime := time.Now().Unix()
//	var apiLogChannel = make(chan model.NewLogData)
//	currentNewTitle = binance.GetCurrentNewTitle()
//	c.StartFetchNewTask(apiLogChannel)
//
//	go func() { // 每分钟清一次 normal 日志
//		for {
//			time.Sleep(time.Second * 10)
//			logger.ClearNormmalLog("mobileApi")
//		}
//	}()
//
//	go func() {
//		c.StartFetchGateData()
//	}()
//
//	for {
//		logData, ok := <-apiLogChannel
//		if !ok || !c.KeepRunning {
//			break
//		}
//		if logData.Success {
//			title := logData.Content
//			titleType := logData.TitleType
//			publishTime := logData.PublishTime
//			logger.Addlog("domain: "+logData.Domain+"；"+title+"——"+titleType, "normal", "mobileApi")
//			isEixst := tool.StringInArray(title, c.ExistTitles)
//			if !isEixst {
//				logger.Addlog("新标题："+title+"；类型："+titleType, "normal", "mobileApi")
//				coins := c.HandleNewTitle(title)                                  // 从标题中提取 coin，可能存在多个
//				find, coin, usdtCoinItem, _, checkMsg := c.HandleGateCoins(coins) // 从 gaste.io 的币中查找这个币是否存在
//				logger.Addlog(checkMsg, "normal", "operate")
//				logger.Addlog("标题发布时间："+publishTime, "normal", "operate")
//				if find && titleType == "New Cryptocurrency Listing" && c.KeepRunning {
//					c.Stop()
//					logger.Addlog("coin = "+coin, "normal", "operate")
//					fmt.Printf("usdtCoinItem = %v \n", usdtCoinItem)
//					//tradeResult := c.DoGateTrade(coin, usdtCoinItem)
//					//c.HandleTradeResult(tradeResult, usdtCoinItem)
//
//					logger.Addlog("更新币数据表", "normal", "operate")
//					// ====================== 添加新币到数据库 start ======================
//					global.DB_ENGINE.Table("bm_exist_coin").First(&db_coin)
//					if db_coin.Coin == "" {
//						db_coin.Coin = coin
//					} else {
//						completeCoinData := db_coin.Coin + "," + coin
//						db_coin.Coin = completeCoinData
//					}
//					global.DB_ENGINE.Table("bm_exist_coin").Save(&db_coin)
//					// ====================== 添加新币到数据库 end ======================
//					c.UpdateBinanceNewTitle(title, titleType, publishTime, 20)
//					break
//				}
//				c.UpdateBinanceNewTitle(title, titleType, publishTime, 10)
//			}
//		} else {
//			logger.Addlog("domain = "+logData.Domain, "error", "error")
//			logger.Addlog(logData.Content, "error", "error")
//			if c.KeepRunning {
//				endTime := time.Now().Unix()
//				logger.Addlog("运行时长："+tool.Int64ToString(endTime-startTime), "error", "error")
//				time.Sleep(time.Minute * 20)
//				logger.Addlog("程序重启", "error", "error")
//				startTime = time.Now().Unix()
//				//go func() {
//				//	c.RunNewTask()
//				//}()
//				//break
//			} else {
//				c.Stop()
//				logger.Addlog("程序结束", "error", "error")
//				continue
//			}
//		}
//	}
//}

func (c *Concurrent) StartFetchGateData() {
	for {
		if !c.KeepRunning {
			break
		}
		gate.SetGateCoinPrice()
		time.Sleep(time.Second * 9)
	}
}

func (c *Concurrent) StartFetchTask(ch chan model.LogData) {
	c.KeepRunning = true
	for i := 0; i < c.WorkerCount; i++ {
		go func() {
			c.Fetch(ch, "mobileApi")
		}()
	}
}

func (c *Concurrent) StartFetchNewTask(ch chan model.NewLogData) {
	c.KeepRunning = true
	for i := 0; i < c.WorkerCount; i++ {
		go func() {
			c.NewFetch(ch, "mobileApi")
		}()
	}
}

func (c *Concurrent) NewFetch(logChannel chan model.NewLogData, source string) {
	for {
		if !c.KeepRunning {
			break
		}
		success := false
		statusCode := 0
		title, titleType, publishTime := "", "", ""
		level := "normal"
		var err error
		// 随机获取域名
		rand.Seed(time.Now().UnixNano() / 1e6)
		index := rand.Intn(len(c.Domains))
		fmt.Printf("index = %v \n", index)
		domain := c.Domains[index]
		success, statusCode, title, titleType, publishTime, err = binance.FetchFromNewApi(domain)
		if err != nil {
			level = "error"
		}
		logData := model.NewLogData{
			Success:     success,
			StatusCode:  statusCode,
			Time:        time.Now().Unix(),
			Content:     title,
			TitleType:   titleType,
			PublishTime: publishTime,
			Level:       level,
			Source:      source,
			Domain:      domain,
		}
		logChannel <- logData
		time.Sleep(time.Duration(c.NewRequestTimeGap) * time.Millisecond)
		//fmt.Printf("title = %v \n", logData.Content)
	}
}

func (c *Concurrent) Fetch(logChannel chan model.LogData, source string) {
	index := 0
	domainIndex := 0
	for {
		if domainIndex == 0 {
			logger.Addlog("开始轮询时间：" + tool.Int64ToString(time.Now().Unix()), "normal", "mobileApi")
		}
		if !c.KeepRunning {
			break
		}
		success := false
		statusCode := 0
		var createTime int64
		title := ""
		level := "normal"
		var err error
		// ===============================================================================
		// 测试代码开始					// 1%的概率让 index 等于 n 作为参数传入 FetchFromApi
		//rand.Seed(time.Now().Unix() + 20)
		//rendResult := rand.Intn(1000)
		////logger.Addlog("rendResult = " + strconv.Itoa(rendResult), "normal", "mobileApi")
		//fmt.Printf("rendResult = %v \n", rendResult)
		//if rendResult <= 10 {
		//	index = 1
		//} else if (rendResult > 10 && rendResult <= 300) {
		//	index = 0
		//} else if (rendResult > 300 && rendResult <= 600) {
		//	index = 2
		//} else if (rendResult > 600 && rendResult <= 1000) {
		//	index = 3
		//}
		// ===============================================================================
		//rand.Seed(time.Now().UnixNano() / 1e6)
		//domainIndex := rand.Intn(len(c.Domains))
		domain := c.Domains[domainIndex]
		success, statusCode, title, createTime, err = binance.FetchFromApi(index, domain)
		if err != nil {
			level = "error"
		}
		logData := model.LogData{
			Success:     success,
			StatusCode:  statusCode,
			Time:        time.Now().Unix(),
			Content:     title,
			Level:       level,
			Source:      source,
			DomainIndex: domainIndex,
			CreateTime:  createTime,
		}
		logChannel <- logData
		//fmt.Printf("title = %v \n", logData.Content)
		if domainIndex < len(c.Domains)-1 {
			domainIndex += 1
		} else {
			time.Sleep(time.Duration(c.RequestTimeGap) * time.Millisecond)
			domainIndex = 0
		}
	}
}

func (c *Concurrent) FetchFromNewListing(logChannel chan model.LogData, source string) {
	for {
		if !c.KeepRunning {
			break
		}
		success := false
		statusCode := 0
		title := ""
		level := "normal"
		var err error
		success, statusCode, title, err = binance.FetchNewListing()
		if err != nil {
			level = "error"
		}
		logData := model.LogData{
			Success:    success,
			StatusCode: statusCode,
			Time:       time.Now().Unix(),
			Content:    title,
			Level:      level,
			Source:     source,
		}
		logChannel <- logData
		time.Sleep(time.Duration(c.RequestTimeGap) * time.Millisecond)
		//fmt.Printf("title = %v \n", logData.Content)
	}
}

func (c *Concurrent) HandleNewTitle(title string) (coins []string) {
	var matchCoins []string
	global.FIND_NEW_TITLE_TIME = time.Now().UnixNano() / 1e6
	logger.Addlog("发现时间： "+strconv.FormatInt(global.FIND_NEW_TITLE_TIME, 10), "normal", "operate")
	//logger.Addlog("步骤一： " + strconv.FormatInt(time.Now().UnixNano() / 1e6,10), "normal", "operate")
	if strings.Contains(title, "Binance Will List") {
		matchCoins = coinReg.FindStringSubmatch(title)
	}
	return matchCoins
}

func (c *Concurrent) HandleGateCoins(coins []string) (bool, string, model.CoinItem, model.CoinItem, string) {
	find := false
	targetCoin := ""
	msg := ""
	var usdtCoinResult model.CoinItem
	var ethCoinResult model.CoinItem
	for _, v := range coins {
		targetCoin = v
		usdtValue, uok := global.GATE_MONITOR_COIN[v+"_USDT"]
		//ethValue, eok := global.GATE_MONITOR_COIN[v+"_ETH"]
		if uok {
			isEixst := tool.StringInArray(targetCoin, c.ExistCoin)
			if !isEixst {
				msg = "找到目标币种：" + targetCoin
				find = true
				usdtCoinResult = usdtValue
				//ethCoinResult = ethValue
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

func (c *Concurrent) DoGateTrade(coin string, coinPriceData model.CoinItem) model.GateOrderResult {
	coinPrice := strconv.FormatFloat(coinPriceData.Price, 'f', coinPriceData.Prec, 64)
	tempPrice := coinPriceData.Price * (1 + c.Slippage/100)
	tempQuantity, _ := decimal.NewFromFloat(c.BuyUsdt / tempPrice).Round(3).Float64()
	quantity := strconv.FormatFloat(tempQuantity, 'f', 2, 64) // 转换 string

	BuyUsdtStr := strconv.FormatFloat(c.BuyUsdt, 'f', 3, 64)
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

func (c *Concurrent) DoEthTrade(coinPriceData model.CoinItem) {
	var buyEth float64
	buyEth = 2
	coinPrice := strconv.FormatFloat(coinPriceData.Price, 'f', coinPriceData.Prec, 64)
	tempPrice := coinPriceData.Price * (1 + c.Slippage/100)
	tempQuantity, _ := decimal.NewFromFloat(buyEth / tempPrice).Round(3).Float64() // 默认以2个ETH购买
	quantity := strconv.FormatFloat(tempQuantity, 'f', 2, 64)
	BuyEthStr := strconv.FormatFloat(buyEth, 'f', 3, 64)
	price := strconv.FormatFloat(tempPrice, 'f', coinPriceData.Prec, 64)
	logMsg := "购买数量：使用eth = " + BuyEthStr + "；数量 = " + quantity + "；价格 = " + price + "；发现价：" + coinPrice
	logger.Addlog(logMsg, "normal", "operate")
}

func (c *Concurrent) HandleTradeResult(tradeRes model.GateOrderResult, coinItem model.CoinItem) {
	if tradeRes.Id != "" {
		if tradeRes.Left == "0" {
			logger.Addlog("Fit!! 全部成交", "normal", "operate")
		} else {
			c.HasOrder = true
			logger.Addlog("部分成交。剩余 "+tradeRes.Left+" 未成交。", "normal", "operate")
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

func (c *Concurrent) UpdateBinanceTitle(title string, operateType int) {
	// 1. 添加到 c.ExistTitles
	// 2. 添加到数据库。先查重，再添加，如果添加成功就发送短信通知
	c.ExistTitles = append(c.ExistTitles, title)
	isAddNew := binance.UpdateNewTitle(title)
	logger.Addlog("isAddNew = "+strconv.FormatBool(isAddNew), "normal", "operate")
	if isAddNew && operateType == 10 {
		logger.Addlog(title, "normal", "operate")
		logger.Addlog("新标题通知", "normal", "operate")
		informer.InformationNewListing(global.NEW_LISTING_INFO_TPL_CODE, 1)
	}
}

func (c *Concurrent) UpdateBinanceNewTitle(title, titleType, publishTime string, operateType int) {
	// 1. 添加到 c.ExistTitles
	// 2. 添加到数据库。先查重，再添加，如果添加成功就发送短信通知
	c.ExistTitles = append(c.ExistTitles, title)
	isAddNew := binance.UpdateNewTitleFromNewApi(title, titleType, publishTime)
	logger.Addlog("isAddNew = "+strconv.FormatBool(isAddNew), "normal", "operate")
	if isAddNew && operateType == 10 {
		logger.Addlog(title, "normal", "operate")
		logger.Addlog("新标题通知", "normal", "operate")
		informer.InformationNewListing(global.NEW_LISTING_INFO_TPL_CODE, 1)
	}
}

func (c *Concurrent) Stop() {
	c.KeepRunning = false
}

func (c *Concurrent) Restart() {
	c.KeepRunning = true
}
