package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"my_e_commerce/data/model"
	"my_e_commerce/service"
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

type UpdateRequest struct {
	oldUser model.User `json:"olduser"`
	user    model.User `json:"user"`
}

func (h *UserHandler) UpdateUserById(c *gin.Context) {
	var updateRequest UpdateRequest
	err := c.BindJSON(&updateRequest)
	if err != nil {
		log.Printf("unmarshal error, %+v", err)
		c.JSON(400, gin.H{"error": "无效的用户数据"})
		return
	}
	fmt.Printf("updateRequest: %+v\n", updateRequest)
	fmt.Printf("oldusrname: %+v\n", updateRequest.oldUser)
	fmt.Printf("usrname: %+v\n", updateRequest.user)
	err = h.userService.UpdateUser(&updateRequest.oldUser, &updateRequest.user)
	if err != nil {
		log.Printf("sqlerror, %+v", err)
		c.JSON(500, gin.H{"error": "用户信息更新失败"})
		return
	}
	c.JSON(200, updateRequest.user)
}
