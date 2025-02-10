package controller

import (
	"github.com/gin-gonic/gin"
	"myproject/forum/server/models"
	"myproject/forum/server/service"
	"net/http"
)

type UserController struct {
	UserService service.IUserService
}

func NewUserController(userService service.IUserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

func (uc *UserController) Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := uc.UserService.CreateUser(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "User created successfully")
}
