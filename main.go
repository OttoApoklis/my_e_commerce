package main

import "github.com/gin-gonic/gin"

func main() {
	// 创建一个gin路由
	r := gin.Default()

	// 定义一个简单的GET路由
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})
	r.Run(":8080")
}
