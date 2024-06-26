package api

import (
	"binanceNewCoin/server/exchanger/binance"
	"binanceNewCoin/server/exchanger/gate"
	"binanceNewCoin/server/exchanger/mx"
	"binanceNewCoin/server/global"
	"binanceNewCoin/server/global/response"
	"binanceNewCoin/server/logger"
	"binanceNewCoin/server/model"
	"binanceNewCoin/server/model/request"
	"binanceNewCoin/server/tool"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var coinReg = regexp.MustCompile(`\((.*?)\)`)
var ExistCoin = []string{"GTC", "KLAY"}
var testData string
var isMain bool

var tags = [5]string{"65uLk7Nkp0pRea2ps4Aj", "Z73vs0d3eu67rDqMCbSH", "tfph2mpTPAuwxbiMHoQc", "pch5D9lsORjgObhyjdSK", "tvXLzOPgJFiMa8Omltoo"}

func TestGateOrder(c *gin.Context) {
	var testData model.TestGateOrder
	_ = c.ShouldBindJSON(&testData)
	usdt, _ := tool.StringToFloat(testData.Usdt,64)
	//slippage, _ := tool.StringToFloat(testData.Slippage,64)
	price, _ := tool.StringToFloat(testData.Price,64)
	coin := testData.Coin
	fmt.Printf("testData = %v \n", testData)
	amount, _ := decimal.NewFromFloat(usdt / price).Round(3).Float64()
	amountString := strconv.FormatFloat(amount, 'f', 2, 64)
	gateOrderData := gate.GetOrderData(coin, testData.Side, amountString, testData.Price, testData.Account, testData.Auto_borrow)
	fmt.Printf("gateOrderData = %v \n", gateOrderData)
	gate.DoTrade(gateOrderData)
	response.OkWithMessage("ok", c)
}

func TestTransfers(c *gin.Context) {
	var transferBody model.TransfersBody
	_ = c.ShouldBindJSON(&transferBody)
	gate.TransfersAccount(transferBody)
	response.OkWithMessage("ok", c)
}

func TestGetMarginCoins(c *gin.Context) {
	gate.GetGateMarginCoins()
	response.OkWithMessage("ok", c)
}

func GetCurrentTitle(c *gin.Context) {
	currentNewTitle := binance.GetCurrentNewTitle()
	data := map[string]string {
		"title": currentNewTitle,
	}
	response.Result(0, data, "", c)
}

func TestMxOrder(c *gin.Context) {
	var mxOrder request.MxOrder
	_ = c.ShouldBindJSON(&mxOrder)
	orderParam := map[string]string{
		"symbol":     mxOrder.Coin + "_USDT",
		"price":      mxOrder.Price,
		"quantity":   mxOrder.Quantity,
		"trade_type": "BID",
		"order_type": "IMMEDIATE_OR_CANCEL",		// 下单后马上撤销，以防部分成交，价格回落时买入
	}
	m := &mx.Mx{
		AccessKey: mxOrder.Access_key,
		SecretKey: mxOrder.Secret_key,
	}
	result, err := m.OrderBuy(orderParam)
	if err != nil {
		fmt.Printf("err = %v \n", err)
		response.Result(-1, err, "下单失败", c)
		return
	}
	fmt.Printf("result = %v \n", result)
	if result["code"] != 200 {
		response.Result(-1, result, "下单失败", c)
		return
	}
	response.Result(0, result, "下单成功", c)
}

func TestGetApi (c *gin.Context) {
	coin := ""
	log := ""
	startTime := time.Now().UnixNano() / 1000000
	logger.Addlog("开始时间：" + strconv.FormatInt(startTime, 10), "normal", "mobileApi")
	success, statusCode, listingTitle, err := binance.TestFetchFromApi(0)
	endRequestTime := time.Now().UnixNano() / 1000000
	logger.Addlog("完成请求时间：" + strconv.FormatInt(endRequestTime, 10), "normal", "mobileApi")
	logger.Addlog("耗时：" + strconv.FormatInt(endRequestTime - startTime, 10), "normal", "mobileApi")
	//fmt.Printf("listingTitle = %v \n", listingTitle)
	//fmt.Printf("err = %v \n", err)
	log = "请求日志：success = " + strconv.FormatBool(success) + "；statusCode = " + strconv.Itoa(statusCode) + "；listingTitle = " + listingTitle
	if err != nil {
		log = log + "；err = " + err.Error()
	}
	logger.Addlog(log, "error", "error")
	if strings.Contains(listingTitle, "Binance Will List") {
		fmt.Printf("含有关键字 \n")
		matchCoin := coinReg.FindStringSubmatch(listingTitle)
		fmt.Printf("matchCoin = %v \n", matchCoin)
		if len(matchCoin) > 1 {
			coin = matchCoin[1]
		}
	}

	//fmt.Printf("coin = %v \n", coin)
	// 查询获取到的币种是否已存在，已存在就不做处理，不存在就发起 exchanger
	isEixst := tool.StringInArray(coin, ExistCoin)
	finishTime := time.Now().UnixNano() / 1000000
	logger.Addlog("完成分析时间：" + strconv.FormatInt(finishTime, 10), "normal", "mobileApi")
	logger.Addlog("最终耗时：" + strconv.FormatInt(finishTime - startTime, 10), "normal", "mobileApi")

	_err := ""
	if err != nil {
		_err = err.Error()
	}
	data := map[string]interface{} {
		"success": success,
		"coin": coin,
		"isEixst": isEixst,
		"err": _err,
	}
	response.Result(0, data, "", c)
}

func TestAddGateOrder(c *gin.Context) {
	o := &model.DB_Gate_Order {
		Order_id: "76939752188",
		Coin: "GALA",
		Price: "0.02494",
		Amount: "120271.81",
		Total: "2925.0104192",
		Left: "0",
		Create_time: "1631498451",
	}
	fmt.Printf("gos = %v \n", o)
	err := global.DB_ENGINE.Table("bm_order").Create(o).Error
	if err != nil {
		response.OkWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("添加数据成功", c)
	}
}

func GetTestData(c *gin.Context) {
	if testData == "true" {
		isMain = true
	}
	result := map[string]interface{} {
		"isMain": isMain,
	}
	response.Result(0, result, "", c)
}

func TestRandNumber(c *gin.Context) {
	rendResult := rand.Intn(1000)
	fmt.Printf("rendResult = %v \n", rendResult)
	response.Ok(c)
}

func CheckBinanceTitle(c *gin.Context) {
	var article request.Article
	var articles []model.Article
	_ = c.ShouldBindJSON(&article)
	newTitle := article.Title
	fmt.Printf("newTitle = %v \n", newTitle)
	global.DB_ENGINE.Table("bm_articles").Where("content = ?", newTitle).Find(&articles)
	fmt.Printf("articles = %v \n", articles)
	fmt.Printf("articles length = %v \n", len(articles))
	if (len(articles) == 0) {
		create_time := strconv.FormatInt(time.Now().UnixNano() / 1e6, 10)
		db_article := model.Article {
			Content: newTitle,
			Create_time: create_time,
		}
		result := global.DB_ENGINE.Table("bm_articles").Create(&db_article)
		fmt.Printf("result %v \n", result)
	} else {
		fmt.Printf("标题已存在 \n")
	}
	response.Ok(c)
}

func GetBinanceAllCoin(c *gin.Context) {
	binanceCoin, err := binance.GetAll()
	if err != nil {
		fmt.Printf("err = %v \n", err)
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithData(binanceCoin, c)
	}
}

func TestBinanceNewApi(c *gin.Context) {
	startTime := time.Now().UnixNano() / 1000000
	logger.Addlog("开始时间：" + strconv.FormatInt(startTime, 10), "normal", "mobileApi")
	binance.TestFetchNewApi()
	endRequestTime := time.Now().UnixNano() / 1000000
	logger.Addlog("完成请求时间：" + strconv.FormatInt(endRequestTime, 10), "normal", "mobileApi")
	logger.Addlog("耗时：" + strconv.FormatInt(endRequestTime - startTime, 10), "normal", "mobileApi")
	response.Ok(c)
}

func TestQuery(c *gin.Context) {
	name := c.Query("name")
	fmt.Printf("name = %v \n", name)
	response.Ok(c)
}

func TestGetRandom(c *gin.Context) {
	rand.Seed(time.Now().Unix())
	result := rand.Intn(5)
	fmt.Printf("result = %v \n ", tool.IntToString(result))
	fmt.Printf("tag = %v \n", tags[result])
	fmt.Printf("====================================== \n")
	response.Ok(c)
}
