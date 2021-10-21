package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"goBlog/utils"
	"time"
)

var db *gorm.DB
var err error

func InitDb() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DbUser,utils.DbPassword, utils.DbHost, utils.DbPort, utils.DbName)
	fmt.Println(dsn)
	db, err = gorm.Open(utils.Db, dsn) // := 的声明方式是一个局部变量
	if err != nil {
		fmt.Printf("数据库连接错误请检查参数！%v/n", err)
	}
	// 禁用复数形式的表名
	db.SingularTable(true)

	// 数据迁移
	db.AutoMigrate(&User{}, &Category{}, &Article{})

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	db.DB().SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	db.DB().SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	db.DB().SetConnMaxLifetime(10 * time.Second) // 不要大于框架的超时时间

	//db.Close()
}
