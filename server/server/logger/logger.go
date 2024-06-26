package logger

import (
	"binanceNewCoin/server/global"
	"binanceNewCoin/server/model"
	"fmt"
	"time"
)

func AddTitlelog (logData model.LogData) {
	//fmt.Printf("source = %v ", logData.Source)
	//fmt.Printf(" title = %v \n", logData.Content)
	//fmt.Printf("======================================================= \n")
	if logData.Level == "normal" {
		if logData.Source == "announce" {
			global.ANNOUNCE_LOGS = append(global.ANNOUNCE_LOGS, logData)
		} else if logData.Source == "newListing" {
			global.NEW_LISTING_LOGS = append(global.NEW_LISTING_LOGS, logData)
		} else if logData.Source == "mobileApi" {
			global.API_LOGS = append(global.API_LOGS, logData)
		}
	} else if logData.Level == "error" {
		fmt.Printf("报错日志：title = %v \n", logData.Content)
		fmt.Printf("报错日志：source = %v \n", logData.Source)
		global.ERROR_LOGS = append(global.ERROR_LOGS, logData)
	}
}

func Addlog (data string, level string, source string) {
	logData := model.LogData{
		Time: time.Now().UnixNano() / 1e6,
		Content: data,
		Level: level,
		Source: source,
	}
	//if level == "normal" {
	//	global.NORMAL_LOGS = append(global.NORMAL_LOGS, logData)
	//} else {
	//	global.LOGS = append(global.LOGS, logData)
	//}
	if source == "operate" {
		global.OPERATE_LOGS = append(global.OPERATE_LOGS, logData)
	} else if source == "error" {
		global.ERROR_LOGS = append(global.ERROR_LOGS, logData)
	} else if source == "mobileApi" {
		global.API_LOGS = append(global.API_LOGS, logData)
	} else if source == "exchange" {
		global.EXCHANGE_LOGS = append(global.EXCHANGE_LOGS, logData)
	}
}

func ClearNormmalLog (source string) {
	// 保留最新20条

	if (len(global.API_LOGS) > 200) && (source == "mobileApi") {
		global.API_LOGS = global.API_LOGS[len(global.API_LOGS) - 200: len(global.API_LOGS)]
	}
	if (len(global.EXCHANGE_LOGS) > 20) && (source == "exchange") {
		global.EXCHANGE_LOGS = global.EXCHANGE_LOGS[: 0]
	}
	if (len(global.ERROR_LOGS) > 20) && (source == "error") {
		global.ERROR_LOGS = global.ERROR_LOGS[: 0]
	}

	//if (len(global.ANNOUNCE_LOGS) > 15) && (source == "announce") {
	//	global.ANNOUNCE_LOGS = global.ANNOUNCE_LOGS[len(global.ANNOUNCE_LOGS) - 15: len(global.ANNOUNCE_LOGS)]
	//}
	//if (len(global.NEW_LISTING_LOGS) > 15) && (source == "newListing") {
	//	global.NEW_LISTING_LOGS = global.NEW_LISTING_LOGS[len(global.NEW_LISTING_LOGS) - 15: len(global.NEW_LISTING_LOGS)]
	//}
	//if (len(global.API_LOGS) > 15) && (source == "mobileApi") {
	//	global.API_LOGS = global.API_LOGS[len(global.API_LOGS) - 15: len(global.API_LOGS)]
	//}
}
