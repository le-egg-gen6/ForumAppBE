package handler

import (
	"forum/shared"
	"github.com/gin-gonic/gin"
)

func InitializePostHandler(router *gin.RouterGroup) {
	postGroup := router.Group("post")
	{
		postGroup.GET("/test", func(c *gin.Context) {
			shared.SendSuccess(c, "Ok!")
		})
	}
}
