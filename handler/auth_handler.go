package handler

import (
	"forum/shared"
	"github.com/gin-gonic/gin"
)

func InitializeAuthHandler(router *gin.RouterGroup) {
	authGroup := router.Group("/auth")
	{
		authGroup.GET("/test", Login)
	}
}

func Login(c *gin.Context) {
	shared.SendSuccess(c, "Test")
}
