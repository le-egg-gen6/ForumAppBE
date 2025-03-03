package controller

import (
	"github.com/gin-gonic/gin"
	"myproject/forum/dtos"
	"myproject/forum/models"
	"myproject/forum/service"
	"myproject/forum/shared"
	"myproject/forum/util"
	"net/http"
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
	var loginDTO dtos.LoginDTO
	if err := c.ShouldBind(&loginDTO); err != nil {
		shared.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := ac.UserService.GetUserByUsername(loginDTO.Username)
	if err != nil {
		shared.SendError(c, http.StatusBadRequest, err.Error())
		return
	}
	if user == nil {
		shared.SendError(c, http.StatusBadRequest, "user not found")
		return
	}

	check := util.CheckPasswordHash(loginDTO.Password, user.Password)
	if !check {
		shared.SendError(c, http.StatusUnauthorized, "invalid password")
		return
	}

	token, err := util.GenerateToken(user.ID, loginDTO.Remember)
	if err != nil {
		shared.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	shared.SendSuccess(c, dtos.AuthDTO{Token: token})
}

func (ac *AuthController) Register(c *gin.Context) {
	var registerDTO dtos.RegisterDTO
	if err := c.ShouldBind(&registerDTO); err != nil {
		shared.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	user_1, err := ac.UserService.GetUserByUsername(registerDTO.Username)
	user_2, err := ac.UserService.GetUserByEmail(registerDTO.Email)
	if err != nil {
		shared.SendError(c, http.StatusBadRequest, err.Error())
		return
	}
	if user_1 != nil || user_2 != nil {
		shared.SendError(c, http.StatusBadRequest, "user is already registered")
		return
	}

	hashedPassword, err := util.HashPassword(registerDTO.Password)
	if err != nil {
		shared.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	user := &models.User{
		Username: registerDTO.Username,
		Email:    registerDTO.Email,
		Password: hashedPassword,
	}

	user, err = ac.UserService.CreateUser(user)
	if err != nil {
		shared.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	token, err := util.GenerateToken(user.ID, false)
	if err != nil {
		shared.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}
	shared.SendSuccess(c, dtos.AuthDTO{Token: token})
}
