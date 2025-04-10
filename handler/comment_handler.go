package handler

import (
	"forum/shared"
	"github.com/gin-gonic/gin"
)

func InitializeCommentHandler(router *gin.RouterGroup) {
	commentGroup := router.Group("/comment")
	{
		commentGroup.GET("/test", func(c *gin.Context) {
			shared.SendSuccess(c, "Ok!")
		})
	}
}
