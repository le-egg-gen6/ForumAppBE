package handler

import (
	"forum/3rd_party_service/mail_sender"
	"forum/dtos"
	"forum/models"
	"forum/repository"
	"forum/shared"
	"forum/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitializeAuthHandler(router *gin.RouterGroup) {
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/login", Login)
		authGroup.POST("/register", Register)
	}
}

func Login(c *gin.Context) {
	var loginRequest dtos.LoginDTO
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		shared.SendError(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	username := loginRequest.Username
	password := loginRequest.Password
	remember := loginRequest.Remember

	user, err := repository.GetUserRepositoryInstance().FindByUsername(username)
	if err != nil {
		shared.SendInternalServerError(c)
		return
	}
	if user == nil {
		shared.SendError(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		shared.SendError(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	token, err := utils.GenerateToken(user.ID, remember)
	if err != nil {
		shared.SendInternalServerError(c)
		return
	}

	shared.SendSuccess(c, dtos.AuthDTO{Token: token, Validated: user.Validated})
}

func Register(c *gin.Context) {
	var registerDTO dtos.RegisterDTO
	if err := c.ShouldBindJSON(&registerDTO); err != nil {
		shared.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	username := registerDTO.Username
	email := registerDTO.Email
	password := registerDTO.Password

	user, err := repository.GetUserRepositoryInstance().FindByUsername(username)
	if err != nil {
		shared.SendInternalServerError(c)
		return
	}
	if user != nil {
		shared.SendError(c, http.StatusBadRequest, "User already exists")
		return
	}

	emailUser, err := repository.GetUserRepositoryInstance().FindByEmail(email)
	if err != nil {
		shared.SendInternalServerError(c)
		return
	}
	if emailUser != nil {
		shared.SendError(c, http.StatusBadRequest, "Email already exists")
		return
	}

	hashedPw, err := utils.HashPassword(password)
	if err != nil {
		shared.SendInternalServerError(c)
		return
	}

	validateCode := utils.GenerateValidateCode(6)
	utils.ExecuteAsync(mail_sender.SendValidateMail, email, username, validateCode)

	user = &models.User{
		Username:     username,
		Email:        email,
		Password:     hashedPw,
		Validated:    false,
		ValidateCode: validateCode,
	}
	user, err = repository.GetUserRepositoryInstance().Create(user)
	if err != nil {
		shared.SendInternalServerError(c)
		return
	}

	token, err := utils.GenerateToken(user.ID, false)
	if err != nil {
		shared.SendInternalServerError(c)
		return
	}

	shared.SendSuccess(c, dtos.AuthDTO{Token: token, Validated: user.Validated})
}
