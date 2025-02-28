package routes

import (
	"github.com/gin-gonic/gin"
	"myproject/forum/controller"
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
