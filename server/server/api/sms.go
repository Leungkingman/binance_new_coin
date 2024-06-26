package api

import (
	"binanceNewCoin/server/global"
	"binanceNewCoin/server/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

func SendSms(c *gin.Context) {
	as := service.AliSMS{
		AccessKeyId: global.ALI_SMS_AccessKeyId,
		AccessSecret: global.ALI_SMS_AccessSecret,
		Sign: global.ALI_SMS_SIGN,
	}
	for _, v := range global.SMS_CONTACT {
		response, err := as.SendSms(global.NEW_LISTING_INFO_TPL_CODE, v)
		if err != nil {
			fmt.Printf("err = %v \n", err)
		}
		fmt.Printf("response = %v \n", response)
	}
}
