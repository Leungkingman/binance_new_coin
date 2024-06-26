package main

import (
	"binanceNewCoin/server/api"
	"binanceNewCoin/server/middleware"
	"binanceNewCoin/server/service"
	"github.com/gin-gonic/gin"
)

func main() {

	// 初始化数据库
	service.InitDatabase()

	r := gin.Default()
	r.Use(middleware.Cors())

	r.POST("user/login", api.Login)
	UserRouter := r.Group("user").
		Use(middleware.JWTAuth())
	{
		UserRouter.GET("getConfig", api.GetConfig)
		UserRouter.POST("updatePassword", api.UpdatePassword)
		UserRouter.POST("updateConfig", api.UpdateConfig)
		UserRouter.POST("updateSecret", api.UpdateSecret)
	}

	CoinRouter := r.Group("coin").
		Use(middleware.JWTAuth())
	{
		CoinRouter.GET("getCoin", api.GetCoin)
		CoinRouter.POST("updateCoin", api.UpdateCoin)
	}

	DomainRouter := r.Group("domain").
		Use(middleware.JWTAuth())
	{
		DomainRouter.GET("getDomain", api.GetDomain)
		DomainRouter.POST("addDomain", api.AddDomain)
		DomainRouter.POST("updateDomain", api.UpdateDomain)
		DomainRouter.POST("deleteDomain", api.DeleteDomain)
	}

	OrderRouter := r.Group("order").
		Use(middleware.JWTAuth())
	{
		OrderRouter.GET("getGateOrders", api.GetGateOrders)
		OrderRouter.POST("updateGateOrder", api.UpdateGateOrder)
	}

	EngineRouter := r.Group("engine").
		Use(middleware.JWTAuth())
	{
		EngineRouter.GET("startEngine", api.StartEngine)
		EngineRouter.GET("stopEngine", api.StopEngine)
		EngineRouter.GET("getEngineRunningStatus", api.GetEngineRunningStatus)
		EngineRouter.GET("getOrders", api.GetOrders)
		EngineRouter.GET("getEngineLog", api.GetEngineLog)
		EngineRouter.GET("clearEngineLog", api.ClearEngineLog)
		EngineRouter.GET("getQueueTaskRunningStatus", api.GetQueueTaskRunningStatus)
		EngineRouter.POST("startMainQueueTask", api.StartMainQueueTask)
	}

	SecondEngineRouter := r.Group("engine")
	{
		SecondEngineRouter.POST("fetchBinanceData", api.FetchBinanceData)
		SecondEngineRouter.GET("stopQueueTaskRunningStatus", api.StopQueueTaskRunningStatus)
		SecondEngineRouter.GET("startMainQueueWithoutPermission", api.StartMainQueueWithoutPermission)
	}

	ExchangeRouter := r.Group("exchange").
		Use(middleware.JWTAuth())
	{
		ExchangeRouter.GET("getExchangeCoin", api.GetExchangeMonitorCoin)
		ExchangeRouter.GET("getBinanceNewTitle", api.GetBinanceNewTitle)
		ExchangeRouter.GET("getInitialInfo", api.GetInitialInfo)
		ExchangeRouter.GET("initialExchangeData", api.InitialExchangeData)			// 初始化
	}

	TestRouter := r.Group("test")
	{
		TestRouter.POST("testMxOrder", api.TestMxOrder)
		TestRouter.POST("testGateOrder", api.TestGateOrder)
		TestRouter.POST("testTransfers", api.TestTransfers)
		TestRouter.GET("testGetApi", api.TestGetApi)
		TestRouter.GET("getCurrentTitle", api.GetCurrentTitle)
		TestRouter.GET("testAddGateOrder", api.TestAddGateOrder)
		TestRouter.GET("getTestData", api.GetTestData)
		TestRouter.POST("checkBinanceTitle", api.CheckBinanceTitle)
		TestRouter.GET("testRandNumber", api.TestRandNumber)
		TestRouter.GET("getBinanceAllCoin", api.GetBinanceAllCoin)
		TestRouter.GET("testBinanceNewApi", api.TestBinanceNewApi)
		TestRouter.GET("testGetMarginCoins", api.TestGetMarginCoins)
		TestRouter.GET("testQuery", api.TestQuery)
		TestRouter.GET("testGetRandom", api.TestGetRandom)
	}


	SmsRouter := r.Group("sms")
	{
		SmsRouter.GET("send", api.SendSms)
	}

	r.Run(":8880")
}

