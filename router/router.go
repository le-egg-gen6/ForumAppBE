package router

import (
	"github.com/gin-gonic/gin"
	"myproject/forum/config"
	"myproject/forum/middlewares"
)

type Router interface {
	RegisterRoutes(router *gin.RouterGroup)
}

func InitializeRouter(cfg *config.Config, routerModules []Router) *gin.Engine {
	router := gin.Default()

	apiPrefix := "/api/" + cfg.API_VERSION

	api := router.Group(
		apiPrefix,
		middlewares.RecoverMiddleware(),
		middlewares.RequestIDMiddleware(),
		middlewares.LoggerMiddleware(),
		middlewares.CORSMiddleware(),
	)

	for _, module := range routerModules {
		module.RegisterRoutes(api)
	}

	return router
}
