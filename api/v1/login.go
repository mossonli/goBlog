package v1

import (
	"github.com/gin-gonic/gin"
	"goBlog/middleware"
	"goBlog/model"
	"goBlog/utils/errmsg"
	"net/http"
)

func Login(ctx *gin.Context)  {
	var data model.User
	ctx.ShouldBindJSON(&data)

	var token string
	var code int

	code = model.CheckLogin(data.Username, data.Password)

	if code == errmsg.SUCCESS{
		token, code = middleware.SetToken(data.Username)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"message": errmsg.GetErrMsg(code),
		"token": token,
	})
}