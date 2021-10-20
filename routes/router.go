package routes

import (
	"github.com/gin-gonic/gin"
	v1 "goBlog/api/v1"
	"goBlog/utils"
)

func InitRouter(){
	gin.SetMode(utils.AppMode)
	r := gin.Default()
	router := r.Group("api/v1")
	{
		// 用户模块的接口
		router.POST("/user/add", v1.AddUser)
		router.GET("/user/user_exist", v1.UserExist)
		router.GET("/user/get_user_list", v1.GetUserList)
		router.PUT("/user/edit/:id", v1.EditUser)
		router.POST("/user/del/:id", v1.DelUser)
		// 分类模块的接口
		// 文章模块的接口
	}

	r.Run(utils.HttpPort)
}