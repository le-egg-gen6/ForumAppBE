package handler

import (
	"forum/dtos"
	"forum/repository"
	"forum/server/http_server/middlewares"
	"forum/shared"
	"forum/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

func InitializeUserHandler(router *gin.RouterGroup) {
	userGroup := router.Group("/user")
	{
		userGroup.GET("/info", middlewares.AuthenticationMiddlewares(), middlewares.AccountValidationMiddlewares(), GetInfo)
		userGroup.GET("/search", middlewares.AuthenticationMiddlewares(), middlewares.AccountValidationMiddlewares(), SearchUser)
	}
}

func GetInfo(c *gin.Context) {
	userID := utils.GetCurrentContextUserID(c)
	if userID < 0 {
		shared.SendUnauthorized(c)
		return
	}

	user, err := repository.GetUserRepositoryInstance().FindByID(uint(userID))
	if err != nil {
		shared.SendInternalServerError(c)
		return
	}
	if user == nil {
		shared.SendUnauthorized(c)
		return
	}
	searchUserIDStr := utils.GetRequestParam(c, "id")
	searchUserID, err := strconv.Atoi(searchUserIDStr)
	if err != nil {
		shared.SendBadRequest(c, "User not exist")
		return
	}
	searchUser, err := repository.GetUserRepositoryInstance().FindByID(uint(searchUserID))
	if err != nil {
		shared.SendInternalServerError(c)
		return
	}
	if searchUser == nil {
		shared.SendBadRequest(c, "User not exist")
		return
	}
	shared.SendSuccess(c, utils.ConvertToUserDTO(searchUser))
}

func SearchUser(c *gin.Context) {
	userID := utils.GetCurrentContextUserID(c)
	if userID < 0 {
		shared.SendUnauthorized(c)
		return
	}

	user, err := repository.GetUserRepositoryInstance().FindByID(uint(userID))
	if err != nil {
		shared.SendInternalServerError(c)
		return
	}
	if user == nil {
		shared.SendUnauthorized(c)
		return
	}
	searchQuery := utils.GetRequestParam(c, "query")
	searchUsers, err := repository.GetUserRepositoryInstance().FindByPartialUsername(searchQuery)
	if err != nil {
		shared.SendInternalServerError(c)
		return
	}
	userDTOs := make([]dtos.SimpleUserDTO, 0)
	for _, user_ := range searchUsers {
		if user_.ID == user.ID {
			continue
		}
		userDTOs = append(userDTOs, *utils.ConvertToSimpleUserDTO(user_))
	}
	shared.SendSuccess(c, userDTOs)
}
