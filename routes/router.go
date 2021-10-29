package routes

import (
	"github.com/gin-gonic/gin"
	v1 "goBlog/api/v1"
	"goBlog/middleware"
	"goBlog/utils"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())

	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		// 用户模块的接口
		auth.POST("/user/add", v1.AddUser)

		auth.PUT("/user/edit/:id", v1.EditUser)
		auth.DELETE("/user/del/:id", v1.DelUser)
		// 分类模块的接口
		auth.POST("/category/add", v1.AddCategory)

		auth.PUT("/category/edit/:id", v1.EditCategory)
		auth.DELETE("/category/edit/:id", v1.DelCategory)
		// 文章模块的接口
		auth.POST("/article/add", v1.AddArticle)

		auth.PUT("/article/edit/:id", v1.EditArt)
		auth.DELETE("/article/del/:id", v1.DeleteArt)
		// 文件上传
		auth.POST("/upload", v1.UpLoad)
	}
	router := r.Group("api/v1")
	{
		router.GET("/user/user_exist", v1.UserExist)
		router.GET("/user/list", v1.GetUserList)
		router.POST("/user/login", v1.Login)

		router.GET("/category/list", v1.GetCategory)

		router.GET("/article/list", v1.GetArt)
		router.GET("/article/category/articles/:id", v1.GetCateArt)
		router.GET("/article/:id/", v1.GetArtInfo)
	}

	r.Run(utils.HttpPort)
}
