package routes

import (
	"github.com/gin-gonic/gin"
	"myproject/forum/controller"
)

type UserRoutes struct {
	UserController *controller.UserController
}

func (r *UserRoutes) RegisterRoutes(router *gin.RouterGroup) {
	user := router.Group("/users")
	{
		user.POST("/create", r.UserController.Register)
	}

}
