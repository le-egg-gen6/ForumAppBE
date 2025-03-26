package routes

import (
	"forum/controller"
	"github.com/gin-gonic/gin"
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
