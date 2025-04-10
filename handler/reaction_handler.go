package handler

import (
	"forum/shared"
	"github.com/gin-gonic/gin"
)

func InitializeReactionHandler(router *gin.RouterGroup) {
	reactionRouter := router.Group("/reaction")
	{
		reactionRouter.GET("/test", func(c *gin.Context) {
			shared.SendSuccess(c, "Ok!")
		})
	}
}
