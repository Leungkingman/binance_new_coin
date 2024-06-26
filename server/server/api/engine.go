package api

import (
	"binanceNewCoin/server/engine"
	"binanceNewCoin/server/exchanger/mx"
	"binanceNewCoin/server/global"
	"binanceNewCoin/server/global/response"
	"binanceNewCoin/server/middleware"
	"binanceNewCoin/server/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

var e *engine.Concurrent
var m *mx.Mx
var articles []model.Article
var db_domains []model.DB_Domain
var db_coin model.Coin

func StartEngine(c *gin.Context) {
	//timeGap := c.Query("timeGap")
	//sleepTime, _ := strconv.Atoi(timeGap)
	//fmt.Printf("e = %v \n", e)
	if e != nil && e.KeepRunning {
		response.Result(0, nil, "引擎已经运行中", c)
		return
	}
	if !global.INIT_EXCHANGE_FINISH {
		response.Result(-1, nil, "请先初始化系统", c)
		return
	}
	// 获取用户信息
	claims := c.MustGet("claims").(*middleware.CustomClaims)
	uid := claims.Uid
	global.DB_ENGINE.Table("bm_user").First(&user, uid)
	buyUsdt, _ := strconv.ParseFloat(user.Buy_usdt, 64)
	slippage, _ := strconv.ParseFloat(user.Slippage, 64)
	Request_time_gap, _ := strconv.Atoi(user.Request_time_gap)
	New_request_time_gap, _ := strconv.Atoi(user.New_request_time_gap)
	workerCount, _ := strconv.Atoi(user.Worker_count)
	// 获取币种信息
	global.DB_ENGINE.Table("bm_exist_coin").First(&coin)
	coins := strings.Split(coin.Coin, ",")
	//fmt.Printf("coins = %v \n", coins)
	// 获取标题
	tempExistArticles := make([]string, 0)
	global.DB_ENGINE.Table("bm_articles").Find(&articles)
	fmt.Printf("articles = %v \n", articles)
	for i := 0; i < len(articles); i++ {
		tempExistArticles = append(tempExistArticles, articles[i].Content)
	}
	// 获取域名
	existDomains := make([]string, 0)
	global.DB_ENGINE.Table("bm_domain").Find(&db_domains)
	for i := 0; i < len(db_domains); i++ {
		existDomains = append(existDomains, db_domains[i].Domain)
	}
	fmt.Printf("existDomains = %v \n", existDomains)
	e = &engine.Concurrent{
		WorkerCount:       workerCount,
		KeepRunning: 	   true,
		ExistCoin:         coins,
		ExistTitles:       tempExistArticles,
		Domains:           existDomains,
		BuyUsdt:           buyUsdt,
		Slippage:          slippage,
		RequestTimeGap:    Request_time_gap,
		NewRequestTimeGap: New_request_time_gap,
		HasOrder:          false,
	}
	go func() {
		//time.Sleep(time.Duration(sleepTime) * time.Minute)
		e.Run()
		//e.RunNewTask()
	}()
	response.Result(0, nil, "引擎启动成功", c)
}

func StopEngine(c *gin.Context) {
	if e != nil && e.KeepRunning {
		e.Stop()
		response.Ok(c)
		return
	}
	response.Result(-1, nil, "引擎未启动", c)
}

func GetEngineRunningStatus(c *gin.Context) {
	//fmt.Printf("global logs = %v \n", global.LOGS)
	keepRunning := false
	if e != nil && e.KeepRunning {
		keepRunning = true
	}
	data := map[string]interface{}{
		"engineRunning": keepRunning,
		"running_log":   global.API_LOGS,
		"error_log":     global.ERROR_LOGS,
		"operate_log":   global.OPERATE_LOGS,
		"exchange_log":  global.EXCHANGE_LOGS,
	}
	response.Result(0, data, "", c)
}

func GetEngineLog(c *gin.Context) {
	running := false
	if e != nil && e.KeepRunning {
		running = true
	}

	result := map[string]interface{}{
		"engineRunning": running,
		"running_log":   global.API_LOGS,
		"error_log":     global.ERROR_LOGS,
		"operate_log":   global.OPERATE_LOGS,
		"exchange_log":  global.EXCHANGE_LOGS,
	}

	response.Result(0, result, "", c)
}

func ClearEngineLog(c *gin.Context) {
	global.API_LOGS = global.API_LOGS[:0]
	global.ERROR_LOGS = global.ERROR_LOGS[:0]
	global.OPERATE_LOGS = global.OPERATE_LOGS[:0]
	global.EXCHANGE_LOGS = global.EXCHANGE_LOGS[:0]
	response.Result(0, nil, "清除完成", c)
}

func GetOrders(c *gin.Context) {
	claims := c.MustGet("claims").(*middleware.CustomClaims)
	uid := claims.Uid
	global.DB_ENGINE.Table("bm_user").First(&user, uid)

	result := []map[string]interface{}{}
	global.DB_ENGINE.Table("bm_mx_order_id").Limit(10).Find(&result)
	//fmt.Printf("result count = %v \n", dbResult.RowsAffected)
	//fmt.Printf("result = %v \n", result)
	//order_id := result[0]["order_id"]
	order_ids := ""
	for k, v := range result {
		order_ids = order_ids + v["order_id"].(string)
		if k != len(result)-1 {
			order_ids = order_ids + ","
		}
	}
	fmt.Printf("order_ids = %v \n", order_ids)
	m1 := mx.Mx{
		AccessKey: user.Mx_access_key,
		SecretKey: user.Mx_secret_key,
	}
	mxOrderRes, err := m1.GetOrderData(order_ids)
	fmt.Printf("mxOrderRes = %v \n", mxOrderRes)
	if err != nil {
		response.Result(-1, nil, err.Error(), c)
		return
	}
	response.Result(0, mxOrderRes.Data, "", c)
}
