package controller

import (
	"log"
	"my_e_commerce/data/model"
	"my_e_commerce/service"

	"github.com/gin-gonic/gin"
)

//	func UserController() {
//		router := gin.Default()
//
//		userService := serviceImpl.UserServiceImpl{}
//
//		router.POST("/users", func(c *gin.Context) {
//			var newUser model.User
//			if err := c.BindJSON(&newUser); err != nil {
//				log.Println(err)
//				c.JSON(400, gin.H{"error": "无效的用户数据"})
//				return
//			}
//			if err := userService.CreateUser(&newUser); err != nil {
//				c.JSON(500, gin.H{"error": "用户创建失败！"})
//			}
//			c.JSON(201, gin.H{"message": "用户创建成功！"})
//		})
//
//		router.Run(":8080")
//	}
type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var newUser model.User
	if err := c.BindJSON(&newUser); err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"error": "无效的用户数据"})
		return
	}
	if err := h.userService.CreateUser(&newUser); err != nil {
		c.JSON(500, gin.H{"error": "用户创建失败！"})
	}
	c.JSON(201, gin.H{"message": "用户创建成功！"})
}

func (h *UserHandler) SelectByUserId(c *gin.Context) {
	var newUser model.User
	if err := c.BindJSON(&newUser); err != nil {
		log.Printf("unmarshal error, %+v", err)
		c.JSON(400, gin.H{"error": "无效的用户数据"})
		return
	}
	users, err := h.userService.GetUserByID(newUser.ID)
	if err != nil {
		log.Printf("sqlerror, %+v", err)
		c.JSON(500, gin.H{"error": "用户获取失败！"})
		return
	}
	if len(users) > 0 {
		log.Printf(users[0].Username)
	}
	c.JSON(200, users)
}

func (h *UserHandler) UpdateUserById(c *gin.Context) {
	var user model.User
	err := c.BindJSON(&user)
	if err != nil {
		log.Printf("unmarshal error, %+v", err)
		c.JSON(400, gin.H{"error": "无效的用户数据"})
		return
	}
	err = h.userService.UpdateUser(&user)
	if err != nil {
		log.Printf("sqlerror, %+v", err)
		c.JSON(500, gin.H{"error": "用户信息更新失败"})
		return
	}
	c.JSON(200, user)
}
