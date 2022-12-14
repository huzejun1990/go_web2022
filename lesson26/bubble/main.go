package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"net/http"
)



func main() {
	// 创建数据库
	// sql: CREATE DATABASE bubble;
	// 连接数据库
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	defer dao.Close() //程序退出，关闭数据库连接
	//模型绑定
	dao.DB.AutoMigrate(&models.Todo{}) // todos

	r := routers.SetupRouter()
	r.Run()
}
