package controller

import (
	"github.com/gin-gonic/gin"
	"myproject/forum/service"
)

type AuthController struct {
	UserService service.IUserService
}

func NewAuthController(userService service.IUserService) *AuthController {
	return &AuthController{
		UserService: userService,
	}
}

func (ac *AuthController) Login(c *gin.Context) {

}

func (ac *AuthController) Register(c *gin.Context) {

}
