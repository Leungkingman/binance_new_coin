package global

import (
	"binanceNewCoin/server/model"
	"gorm.io/gorm"
)

var DB_ENGINE *gorm.DB
var LOGS []model.LogData
var NORMAL_LOGS []model.LogData

var ANNOUNCE_LOGS []model.LogData
var NEW_LISTING_LOGS []model.LogData
var API_LOGS []model.LogData
var ERROR_LOGS []model.LogData
var OPERATE_LOGS []model.LogData
var EXCHANGE_LOGS []model.LogData

const ALI_SMS_AccessKeyId = "ccccc"
const ALI_SMS_AccessSecret = "ccccccc"
const ALI_SMS_SIGN = "xxx科技"
const NEW_LISTING_INFO_TPL_CODE = "xxxxx"
const GATE_ORDER_TPL_CODE = "xxxxxxx"

var SMS_CONTACT = [...]string{"131cccc"}
