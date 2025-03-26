package routes

import (
	"forum/controller"
	"github.com/gin-gonic/gin"
)

type AuthRoutes struct {
	AuthController *controller.AuthController
}

func (r *AuthRoutes) RegisterRoutes(router *gin.RouterGroup) {
	auth := router.Group("/auth")
	{
		auth.POST("/create")
	}

}
