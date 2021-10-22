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
		router.GET("/user/list", v1.GetUserList)
		router.PUT("/user/edit/:id", v1.EditUser)
		router.DELETE("/user/del/:id", v1.DelUser)
		// 分类模块的接口
		router.POST("/category/add", v1.AddCategory)
		router.GET("/category/list", v1.GetCategory)
		router.PUT("/category/edit/:id", v1.EditCategory)
		router.DELETE("/category/edit/:id", v1.DelCategory)
		// 文章模块的接口
		router.POST("/article/add", v1.AddArticle)
		router.GET("/article/list", v1.GetArt)
		router.GET("/article/category/articles/:id", v1.GetCateArt)
		router.GET("/article/:id/", v1.GetArtInfo)
		router.PUT("/article/edit/:id", v1.EditArt)
		router.DELETE("/article/del/:id", v1.DeleteArt)
	}

	r.Run(utils.HttpPort)
}