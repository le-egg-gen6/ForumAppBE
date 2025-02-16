package router

import (
	"github.com/gin-gonic/gin"
	"myproject/forum/server/config"
)

type Router interface {
	RegisterRoutes(router *gin.RouterGroup)
}

func InitializeRouter(cfg *config.Config, routerModules []Router) *gin.Engine {
	router := gin.Default()

	apiPrefix := "/api/" + cfg.API_VERSION

	api := router.Group(apiPrefix)

	for _, module := range routerModules {
		module.RegisterRoutes(api)
	}

	return router
}
