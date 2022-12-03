package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 定義一個路由，用於處理GET請求
	r.GET("/hello", func(c *gin.Context) {
		// 從請求參數中獲取名稱
		name := c.Query("name")
		// 向客戶端返回一個Hello World!字符串
		c.String(200, "Hello, %s!", name)
	})

	// 啟動Web服務，監聽在8080端口
	r.Run(":8080")
}
