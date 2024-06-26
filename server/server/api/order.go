package api

import (
	"binanceNewCoin/server/exchanger/gate"
	"binanceNewCoin/server/global"
	"binanceNewCoin/server/global/response"
	"binanceNewCoin/server/logger"
	"binanceNewCoin/server/model"
	"binanceNewCoin/server/model/request"
	"fmt"
	"github.com/gin-gonic/gin"
)

var gateOrders []model.DB_Gate_Order

func GetGateOrders(c *gin.Context) {
	global.DB_ENGINE.Table("bm_order").Find(&gateOrders)
	response.Result(0, &gateOrders,"", c)
}

func UpdateGateOrder(c *gin.Context) {
	var orderDetail request.OrderDetailStruct
	_ = c.ShouldBindJSON(&orderDetail)
	orderId := orderDetail.Order_id
	currencyPair := orderDetail.Currency_pair
	gateOrderRes := gate.FetchGateOrder(orderId, currencyPair)
	if gateOrderRes.Id != "" {
		global.DB_ENGINE.Table("bm_order").Where("order_id = ?", orderId).Find(&gateOrders)
		gos, makeErr := gate.MakeUpdateOrderStruct(gateOrderRes)
		fmt.Printf("makeErr = %v \n", makeErr)
		fmt.Printf("gos = %v \n", gos)
		if makeErr == nil {
			if len(gateOrders) == 0 {		//  数据库中没有该订单
				err := global.DB_ENGINE.Table("bm_order").Create(gos).Error
				if err != nil {
					fmt.Printf("err = %v \n", err)
					logger.Addlog("添加订单结果到数据库失败：" + err.Error(), "normal", "operate")
				}
				response.OkWithMessage("添加订单成功", c)
			} else {
				fmt.Printf("gateOrders = %v \n", gateOrders)
				orderItem := gateOrders[0]
				orderItem.Left = gos.Left
				orderItem.Price = gos.Price
				orderItem.Total = gos.Total
				orderItem.Amount = gos.Amount
				global.DB_ENGINE.Table("bm_order").Save(&orderItem)
				response.OkWithMessage("更新订单成功", c)
			}
		} else {
			response.FailWithMessage("构建订单对象失败：" + makeErr.Error(), c)
		}
	} else {
		response.FailWithMessage("找不到订单", c)
	}
}