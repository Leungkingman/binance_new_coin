package tool

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"net/url"
	"strconv"
)

func StringInArray(value string, array []string) bool {
	valueExist := false
	for _, v := range array {
		if v == value {
			valueExist = true
			break
		}
	}

	return valueExist
}

func InArray(value interface{}, array []interface{}) bool {
	valueExist := false
	for _, v := range array {
		if v == value {
			valueExist = true
			break
		}
	}

	return valueExist
}

func Md5Hash(str string) string  {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func ParseRequestParams(params interface{}) (string, *bytes.Reader, error) {
	if params == nil {
		return "", nil, errors.New("illegal parameter")
	}
	data, err := json.Marshal(params)
	if err != nil {
		return "", nil, errors.New("json convert string error")
	}
	jsonBody := string(data)
	binBody := bytes.NewReader(data)
	return jsonBody, binBody, nil
}

func BuildOrderParams(params map[string]string) string {
	urlParams := url.Values{}
	for k := range params {
		urlParams.Add(k, params[k])
	}
	return urlParams.Encode()
}

func CreateSignature(method, path, paramsStr, secret string) string {
	str := method + "\n" + path + "\n" + paramsStr
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func IntToString(number int) string {
	return strconv.Itoa(number)
}

func Int64ToString(number int64) string {
	return strconv.FormatInt(number,10)
}

func StringToFloat(value string, bitSize int) (float64, error) {
	return strconv.ParseFloat(value, bitSize)
}