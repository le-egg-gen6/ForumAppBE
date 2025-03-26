package routes

import (
	"forum/controller"
	"github.com/gin-gonic/gin"
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
