package http

import (
	"context"
	"fmt"
	middlewares2 "forum/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type HTTPServer struct {
	Config      *HTTPServerConfig
	Router      *gin.Engine
	RouterGroup *gin.RouterGroup
	server      *http.Server
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

	httpServer := &http.Server{
		Addr:           fmt.Sprintf(":%d", cfg.Port),
		Handler:        router,
		ReadTimeout:    time.Duration(cfg.ReadTimeoutSec) * time.Second,
		WriteTimeout:   time.Duration(cfg.WriteTimeoutSec) * time.Second,
		MaxHeaderBytes: 1 << 20, // 1 MB
	}

	Server = &HTTPServer{
		Config:      cfg,
		Router:      router,
		RouterGroup: routerGroup,
		server:      httpServer,
	}
}

func GetHTTPServer() *HTTPServer {
	return Server
}

func (httpServer *HTTPServer) Run() error {
	return httpServer.server.ListenAndServe()
}

func (httpServer *HTTPServer) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(httpServer.Config.ShutdownTimeoutSec)*time.Second)
	defer cancel()

	return httpServer.server.Shutdown(ctx)
}
