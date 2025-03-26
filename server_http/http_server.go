package server_http

import (
	"fmt"
	middlewares2 "forum/middlewares"
	"github.com/gin-gonic/gin"
)

type HTTPServer struct {
	Config      *HTTPServerConfig
	Router      *gin.Engine
	RouterGroup *gin.RouterGroup
}

var Server *HTTPServer

func InitializeHTTPServer() {
	cfg, err := LoadHTTPServerConfig()
	if err != nil {
		panic("HTTP Server configuration not found")
	}

	router := gin.Default()

	apiPrefix := "/api/" + cfg.APIVersion

	routerGroup := router.Group(
		apiPrefix,
		middlewares2.RecoverMiddleware(),
		middlewares2.RequestIDMiddleware(),
		middlewares2.LoggerMiddleware(),
		middlewares2.CORSMiddleware(),
	)

	Server = &HTTPServer{
		Config:      cfg,
		Router:      router,
		RouterGroup: routerGroup,
	}
}

func GetHTTPServer() *HTTPServer {
	return Server
}

func (httpServer *HTTPServer) Run() error {
	return httpServer.Router.Run(fmt.Sprintf(":%d", httpServer.Config.Port))
}
