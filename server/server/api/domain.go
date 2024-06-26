package api

import (
	"binanceNewCoin/server/global"
	"binanceNewCoin/server/global/response"
	"binanceNewCoin/server/model"
	"binanceNewCoin/server/model/request"
	"binanceNewCoin/server/service"
	"github.com/gin-gonic/gin"
)

var domains []model.DB_Domain

func GetDomain(c *gin.Context) {
	service.PrintDB()
	global.DB_ENGINE.Table("bm_domain").Find(&domains)
	response.Result(0, &domains,"", c)
}

func AddDomain(c *gin.Context) {
	var addDomain request.AddDomain
	_ = c.ShouldBindJSON(&addDomain)
	db_domain := model.DB_Domain {
		Domain: addDomain.Domain,
	}
	global.DB_ENGINE.Table("bm_domain").Where("domain = ?", addDomain.Domain).Find(&domains)
	//fmt.Printf("domains = %v \n", domains)
	//fmt.Printf("domains len = %v \n", len(domains))
	if len(domains) == 0 {
		global.DB_ENGINE.Table("bm_domain").Create(&db_domain)
		response.Ok(c)
	} else {
		response.FailWithMessage("域名已存在", c)
	}
}

func UpdateDomain(c *gin.Context) {
	var updateDomain request.UpdateDomain
	_ = c.ShouldBindJSON(&updateDomain)
	//fmt.Printf("updateDomain = %v \n", updateDomain)
	global.DB_ENGINE.Table("bm_domain").Where("domain = ?", updateDomain.Domain).Find(&domains)
	if len(domains) == 0 {
		global.DB_ENGINE.Table("bm_domain").Save(&updateDomain)
		response.Ok(c)
	} else {
		response.FailWithMessage("域名已存在", c)
	}
}

func DeleteDomain(c *gin.Context) {
	var deleteDomain request.DeleteDomain
	_ = c.ShouldBindJSON(&deleteDomain)
	//fmt.Printf("updateDomain = %v \n", deleteDomain)
	global.DB_ENGINE.Table("bm_domain").Where("id = ?", deleteDomain.Id).Delete(&model.DB_Domain{})
	response.Ok(c)
}
