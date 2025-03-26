package routes

import (
	"forum/controller"
	"github.com/gin-gonic/gin"
)

type PostRoutes struct {
	PostController *controller.PostController
}

func (r *PostRoutes) RegisterRoutes(router *gin.RouterGroup) {
	posts := router.Group("/posts")
	{
		posts.GET("/test")
	}

}
