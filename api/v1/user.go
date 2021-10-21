package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goBlog/model"
	"goBlog/utils/errmsg"
	"net/http"
	"strconv"
)

type UserInfo struct{

}

var code int


func UserExist(ctx *gin.Context){

}

// 添加用户
func AddUser(ctx *gin.Context){
	var data model.User
	_ = ctx.ShouldBindJSON(&data)
	code = model.CheckUser(data.Username)
	if code == errmsg.SUCCESS{
		model.CreateUser(&data)
	}
	if code == errmsg.ERROR_USERNAME_USED{
		code = errmsg.ERROR_USERNAME_USED
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":code,
		"data": data,
		"message":errmsg.GetErrMsg(code),
	})

}
// 查询用户列表
func GetUserList(ctx *gin.Context){
	pageSize, _ := strconv.Atoi(ctx.Query("pagesize"))
	pageNum, _ := strconv.Atoi(ctx.Query("pagenum"))
	if pageSize == 0{
		pageSize = -1
	}
	if pageNum == 0{
		pageNum = -1
	}
	data := model.GetUserList(pageSize, pageNum)
	code = errmsg.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"data": data,
		"message": errmsg.GetErrMsg(code),
	})
}

// 编辑用户
func EditUser(ctx *gin.Context){
	var data model.User
	id, _ := strconv.Atoi(ctx.Param("id"))
	ctx.ShouldBindJSON(&data)
	code := model.CheckUser(data.Username)
	if code == errmsg.SUCCESS{
		model.EditUser(id, &data)
	}
	if code == errmsg.ERROR_USERNAME_USED {
		ctx.Abort()
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 删除用户
func DelUser(ctx *gin.Context){
	id,_ := strconv.Atoi(ctx.Param("id"))
	fmt.Println("userID", id)
	code := model.DelUser(id)
	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"message": errmsg.GetErrMsg(code),
	})
}
