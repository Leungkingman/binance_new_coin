package api

import (
	"binanceNewCoin/server/global"
	"binanceNewCoin/server/global/response"
	"binanceNewCoin/server/model"
	"binanceNewCoin/server/model/request"
	"binanceNewCoin/server/service"
	"github.com/gin-gonic/gin"
)

var coin model.Coin

func GetCoin(c *gin.Context) {
	service.PrintDB()
	global.DB_ENGINE.Table("bm_exist_coin").First(&coin)
	response.Result(0, &coin,"", c)
}

func UpdateCoin(c *gin.Context) {
	var updateCoins request.UpdateCoinStruct
	_ = c.ShouldBindJSON(&updateCoins)
	global.DB_ENGINE.Table("bm_exist_coin").First(&coin)
	coin.Coin = updateCoins.Coins
	global.DB_ENGINE.Table("bm_exist_coin").Save(&coin)
	response.Ok(c)
}
