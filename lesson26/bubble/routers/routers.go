// @Author huzejun 2022/12/14 19:43:00 
package routers

import (
	"github.com/gin-gonic/gin"
	"huzejun_go/go_web2022/lesson26/bubble/controller"
)

func SetupRouter() *gin.Engine {

	r := gin.Default()
	//告诉gin框架模板文件引用的静态文件支哪里找
	r.Static("/static", "static")

	//告诉gin框架去那里找模板文件
	r.LoadHTMLGlob("templates/*")
	/*	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})*/
	r.GET("/", controller.IndexHandler)

	// v1
	v1Group := r.Group("v1")
	{
		// 待办事项
		// 添加
		v1Group.POST("/todo", controller.CreateTodo)
		// 查看所有的待办事项
		v1Group.GET("/todo", controller.GetTodoList)
		// 修改 某一个待办事项
		v1Group.PUT("/todo/:id", controller.UpdateATodo)
		// 删除 某一个待办事项
		v1Group.DELETE("/todo/:id", controller.DeleteATodo)

	}
	return r
	
}
