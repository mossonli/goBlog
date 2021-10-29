package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"goBlog/utils"
	"goBlog/utils/errmsg"
	"net/http"
	"strings"
	"time"
)

var JwtKey = []byte(utils.JwtKey)
type MyClaims struct {
	Username string  `json:"username"`
	jwt.StandardClaims
}
// 生成token
func SetToken(username string) (string, int) {
	expireTime := time.Now().Add(10*time.Hour)
	setClaims := MyClaims{
		Username:       username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer: "goBlog",
		},
	}
	reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, setClaims)
	token,err := reqClaim.SignedString(JwtKey)
	if err != nil{
		return "", errmsg.ERROR
	}
	return token, errmsg.SUCCESS

}

// 验证token
func CheckToken(token string) (*MyClaims, int) {
	setToken, _:= jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if key,_ := setToken.Claims.(*MyClaims); setToken.Valid{
		return key,errmsg.SUCCESS
	}else {
		return nil, errmsg.ERROR
	}
}

// jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenHeader := ctx.Request.Header.Get("Authorization")
		fmt.Println("token", tokenHeader)
		code := errmsg.SUCCESS
		if tokenHeader == ""{
			code = errmsg.ERROR_TOKEN_NOT_EXIST
			ctx.JSON(http.StatusOK, gin.H{
				"code": code,
				"message": errmsg.GetErrMsg(code),
			})
			ctx.Abort()
			return
		}
		checkToken := strings.SplitN(tokenHeader, " ", 2)
		if len(checkToken) != 2 && checkToken[0] != "Bearer"{
			code = errmsg.ERROR_TOKEN_TYPE_WRONG
			ctx.JSON(http.StatusOK, gin.H{
				"code": code,
				"message": errmsg.GetErrMsg(code),
			})
			ctx.Abort()
			return
		}
		key, tCode := CheckToken(checkToken[1])
		if tCode == errmsg.ERROR{
			code = errmsg.ERROR_TOKEN_WRONG
			ctx.JSON(http.StatusOK, gin.H{
				"code": code,
				"message": errmsg.GetErrMsg(code),
			})
			ctx.Abort()
			return
		}
		if time.Now().Unix() > key.ExpiresAt{
			code = errmsg.ERROR_TOKNE_RUNTIME
			ctx.JSON(http.StatusOK, gin.H{
				"code": code,
				"message": errmsg.GetErrMsg(code),
			})
			ctx.Abort()
			return
		}

		ctx.Set("username", key.Username)
		ctx.Next()
	}
}