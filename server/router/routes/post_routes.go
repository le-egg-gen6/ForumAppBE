package routes

import (
	"github.com/gin-gonic/gin"
	"myproject/forum/server/controller"
)

type PostRoutes struct {
	PostController *controller.PostController
}

func (r *PostRoutes) RegisterRoutes(router *gin.RouterGroup) {
	posts := router.Group("/users")
	{
		posts.GET("/test")
	}

}
