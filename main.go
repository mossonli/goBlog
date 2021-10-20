package main

import (
	"goBlog/model"
	"goBlog/routes"
)

func main() {
	// 初始化数据库
	model.InitDb()

	routes.InitRouter()
}