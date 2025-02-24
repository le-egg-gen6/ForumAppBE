package routes

import (
	"github.com/gin-gonic/gin"
	"myproject/forum/controller"
)

type CommentRoutes struct {
	CommentController *controller.CommentController
}

func (r *CommentRoutes) RegisterRoutes(router *gin.RouterGroup) {
	comments := router.Group("/comments")
	{
		comments.GET("/test")
	}
}
