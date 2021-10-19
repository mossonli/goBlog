package routes

import (
	"github.com/gin-gonic/gin"
	"goBlog/utils"
	"net/http"
)

func InitRouter(){
	gin.SetMode(utils.AppMode)
	r := gin.Default()
	router := r.Group("api/v1")
	{
		router.GET("/hello", func(ctx *gin.Context){
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "ok",
			})
		})
	}

	r.Run(utils.HttpPort)
}