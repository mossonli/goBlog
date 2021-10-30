package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func Cors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cors.New(cors.Config{
			// AllowAllOrigins:        true,	// 和  AllowOrigins 一样的意思 两个保留任意一个
			AllowOrigins:           []string{"*"},
			AllowMethods:           []string{"*"},
			AllowHeaders:           []string{"Origin"},
			ExposeHeaders:          []string{"Content-Length", "Authorization"},
			// AllowCredentials:       true, // 是不是要发送cookie请求
			//AllowOriginFunc: func(origin string) bool {
			//	return origin == "https://github.com"
			//}, // 如果允许所有的域名跨域就可以不要
			MaxAge: 12 * time.Hour,
		})
	}
}
