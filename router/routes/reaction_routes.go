package routes

import (
	"github.com/gin-gonic/gin"
	"myproject/forum/controller"
)

type ReactionRoutes struct {
	ReactionController *controller.ReactionController
}

func (r *ReactionRoutes) RegisterRoutes(router *gin.RouterGroup) {
	posts := router.Group("/reactions")
	{
		posts.GET("/test")
	}

}
