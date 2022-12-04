package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

func main() {
	r := gin.Default()
	//通过multipart forms提交文件时默认的内存限制是32 MiB
	//可以通过下面的方式修改
	//r.MaxMultipartMemory = 8 << 20 // 8 MiB
	r.LoadHTMLFiles("./index.html")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.POST("/upload", func(c *gin.Context) {
		// 从请求中读取文件
		f, err := c.FormFile("f1") // 从请求中获取携带的参数 一样的
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			// 将读取到的文件保存在本地（服务端本地）
			//dst := fmt.Sprintf("./%s",f.Filename)
			dst := path.Join("./", f.Filename)
			c.SaveUploadedFile(f, dst)
			c.JSON(http.StatusOK, gin.H{
				"status": "OK",
			})
		}
		// 将读取到的文件保存在本地（服务端本地）
	})

	r.Run(":8080")
}
