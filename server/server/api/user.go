package api

import (
	"binanceNewCoin/server/global"
	"binanceNewCoin/server/global/response"
	"binanceNewCoin/server/middleware"
	"binanceNewCoin/server/model"
	"binanceNewCoin/server/model/request"
	"binanceNewCoin/server/service"
	"binanceNewCoin/server/tool"
	"fmt"
	"github.com/gin-gonic/gin"
)

type JWT struct {
	SigningKey []byte
}

var user model.User
var users [] model.User

func Login(c *gin.Context) {
	service.PrintDB()
	apiResult := map[string]string{}
	var L request.LoginStruct
	_ = c.ShouldBindJSON(&L)
	username := L.Username
	password := L.Password
	fmt.Printf("username = %v \n", username)
	fmt.Printf("password = %v \n", password)
	if username == "" || password == "" {
		response.Result(-1, nil, "用户名或密码不能为空", c)
		return
	}
	encryptPassword := tool.Md5Hash(password)
	result := global.DB_ENGINE.Table("bm_user").Where(map[string]interface{}{"username": username, "password": encryptPassword}).Find(&users)
	if result.Error != nil {
		response.Result(-1, nil, result.Error.Error(), c)
		return
	}
	if result.RowsAffected > 0 {
		id := &users[0].Id
		username := &users[0].Username
		//fmt.Printf("id = %v \n", *id)
		//fmt.Printf("username = %v \n", *username)
		j := middleware.NewJWT()
		token, err := j.CreateToken(*id, *username)
		if err != nil {
			response.Result(-1, &users, err.Error(), c)
			return
		}
		apiResult["token"] = token
		fmt.Printf("token = %v \n", token)
		response.Result(0, apiResult, "", c)
	} else {
		response.Result(-1, nil, "用户名或密码不正确", c)
	}

}

func GetConfig(c *gin.Context) {
	claims := c.MustGet("claims").(*middleware.CustomClaims)
	uid := claims.Uid
	global.DB_ENGINE.Table("bm_user").First(&user, uid)
	//fmt.Printf("access_key = %v \n", user.Mx_access_key)
	//fmt.Printf("secret_key = %v \n", user.Mx_secret_key)
	if len(user.Mx_access_key) > 0 {
		forntAcsKey := user.Mx_access_key[0:4]
		endAcsKey := user.Mx_access_key[len(user.Mx_access_key)-4 : len(user.Mx_access_key)]
		user.Mx_access_key = forntAcsKey + "************" + endAcsKey
	}
	if len(user.Mx_secret_key) > 0 {
		forntSecKey := user.Mx_secret_key[0:6]
		endSecKey := user.Mx_secret_key[len(user.Mx_secret_key)-6 : len(user.Mx_secret_key)]
		user.Mx_secret_key = forntSecKey + "************" + endSecKey
	}
	response.Result(0, user, "", c)
}

func UpdatePassword(c *gin.Context) {
	claims := c.MustGet("claims").(*middleware.CustomClaims)
	uid := claims.Uid
	var userpass model.Userpassword
	global.DB_ENGINE.Table("bm_user").First(&userpass, uid)
	password := c.PostForm("password")
	encryptPassword := tool.Md5Hash(password)
	userpass.Password = encryptPassword
	global.DB_ENGINE.Table("bm_user").Save(&userpass)
	response.Ok(c)
}

func UpdateConfig(c *gin.Context) {
	claims := c.MustGet("claims").(*middleware.CustomClaims)
	uid := claims.Uid
	global.DB_ENGINE.Table("bm_user").First(&user, uid)

	var U request.UserInfoStruct
	_ = c.ShouldBindJSON(&U)

	user.Buy_usdt = U.Buy_usdt
	user.Min_buy_usdt = U.Min_buy_usdt
	user.Slippage = U.Slippage
	user.Next_slippage = U.Next_slippage
	user.Queue_time_gap = U.Queue_time_gap
	user.Server_queue_time_gap = U.Server_queue_time_gap
	user.Request_time_gap = U.Request_time_gap
	user.New_request_time_gap = U.New_request_time_gap
	user.Long_profit = U.Long_profit
	user.Short_profit = U.Short_profit

	global.DB_ENGINE.Table("bm_user").Save(&user)
	response.Ok(c)
}

func UpdateSecret(c *gin.Context) {
	claims := c.MustGet("claims").(*middleware.CustomClaims)
	uid := claims.Uid
	global.DB_ENGINE.Table("bm_user").First(&user, uid)

	var U request.UserSecretStruct
	_ = c.ShouldBindJSON(&U)

	user.Mx_access_key = U.Mx_access_key
	user.Mx_secret_key = U.Mx_secret_key

	global.DB_ENGINE.Table("bm_user").Save(&user)
	response.Ok(c)
}
