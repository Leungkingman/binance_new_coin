package informer

import (
	"binanceNewCoin/server/global"
	"binanceNewCoin/server/logger"
	"binanceNewCoin/server/service"
)

func InformationNewListing(tplCode string, infoTime int) {
	as := service.AliSMS{
		AccessKeyId:  global.ALI_SMS_AccessKeyId,
		AccessSecret: global.ALI_SMS_AccessSecret,
		Sign:         global.ALI_SMS_SIGN,
	}
	for i := 0; i < infoTime; i++ { // 暂时写死5
		for _, v := range global.SMS_CONTACT {
			response, err := as.SendSms(tplCode, v)
			if err != nil {
				logger.Addlog(err.Error(), "error", "error")
			}
			if response.Code != "OK" {
				errMsg := "发送短信失败，Code = " + response.Code + "；Message = " + response.Message
				logger.Addlog(errMsg, "error", "error")
			}
		}
	}
}
