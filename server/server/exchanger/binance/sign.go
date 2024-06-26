package binance

import (
	"binanceNewCoin/server/global"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func CreateSignature (queryString string) string {
	h := hmac.New(sha256.New, []byte(global.BINANCE_SECRET_KEY))
	h.Write([]byte(queryString))
	return hex.EncodeToString(h.Sum(nil))
}

func GetCompleteQueryString (params map[string]string) string {
	str := ""
	for k, v := range params {
		str += k + "=" + v + "&"
	}
	queryString := str[0:len(str) - 1]
	//fmt.Printf("queryString = %v \n", queryString)
	signature := CreateSignature(queryString)
	return queryString + "&signature=" + signature
}
