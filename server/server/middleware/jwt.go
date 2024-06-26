package middleware

import (
	"binanceNewCoin/server/global/response"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

const JWT_SIGN_KEY = "binancegZu2(b1I0-Dc&3t6mx"

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

type JWT struct {
	SigningKey []byte
}

type CustomClaims struct {
	Uid      int    `json:"uid"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(JWT_SIGN_KEY),
	}
}

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 我们这里jwt鉴权取头部信息 x-token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localSstorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		token := c.Request.Header.Get("x-token")
		if token == "" {
			response.Result(response.ERROR, gin.H{
				"reload": true,
			}, "未登录或非法访问", c)
			c.Abort()
			return
		}

		j := NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == TokenExpired {
				response.Result(response.ERROR, gin.H{
					"reload": true,
				}, "授权已过期", c)
				c.Abort()
				return
			}
			response.Result(response.ERROR, gin.H{
				"reload": true,
			}, err.Error(), c)
			c.Abort()
			return
		}
		//fmt.Printf("claims = %v \n", claims)
		c.Set("claims", claims)
		c.Next()
	}
}

func (j *JWT) CreateToken(uid int, username string) (string, error) {
	maxAge := 60 * 60 * 24		 // 1天
	//maxAge := 10
	customClaims := CustomClaims{
		Uid:      uid,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(maxAge) * time.Second).Unix(), // 过期时间，必须设置
		},
	}
	//采用HMAC SHA256加密算法
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims)
	tokenString, err := token.SignedString([]byte(JWT_SIGN_KEY))
	return tokenString, err
}

func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(JWT_SIGN_KEY), nil
	})
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
