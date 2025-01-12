package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"my_e_commerce/data/model"
	"my_e_commerce/service/serviceImpl"
)

func UserController() {
	router := gin.Default()

	userService := serviceImpl.UserServiceImpl{}

	router.POST("/users", func(c *gin.Context) {
		var newUser model.User
		if err := c.BindJSON(&newUser); err != nil {
			log.Println(err)
			c.JSON(400, gin.H{"error": "无效的用户数据"})
			return
		}
		if err := userService.CreateUser(&newUser); err != nil {
			c.JSON(500, gin.H{"error": "用户创建失败！"})
		}
		c.JSON(201, gin.H{"message": "用户创建成功！"})
	})

	router.Run(":8080")
}
