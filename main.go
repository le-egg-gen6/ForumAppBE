package main

import (
	"fmt"
	"myproject/forum/config"
	"myproject/forum/di"
	"myproject/forum/logger"
	"myproject/forum/router"
	"myproject/forum/tcp"
	"net/http"
)

func main() {
	logger.InitializeLogger()
	logger.GetInstance().Info("================================================================")
	logger.GetInstance().Info("=====================Application Starting=======================")
	logger.GetInstance().Info("================================================================")
	defer CleanupUnfinishedTasks()

	cfg := config.LoadConfig()

	container := di.InitializeContainer(cfg)

	routesModules := []router.Router{
		container.AuthRoutes,
		container.UserRoutes,
		container.PostRoutes,
		container.CommentRoutes,
		container.ReactionRoutes,
	}

	initializeRouter := router.InitializeRouter(cfg, routesModules)
	go func() {
		logger.GetInstance().Info(fmt.Sprintf("HTTP Server running on port: %d", cfg.PORT))
		if err := initializeRouter.Run(fmt.Sprintf(":%d", cfg.PORT)); err != nil {
			logger.GetInstance().Error(fmt.Sprintf("Error starting HTTP server: %s", err))
			panic("Error starting HTTP server")
		}
	}()

	socketServer := tcp.NewSocketServer()
	go func() {
		http.Handle("/socket.io/", socketServer)
		logger.GetInstance().Info(fmt.Sprintf("Socket.IO server running on port: %d", cfg.TCP_PORT))
		if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.TCP_PORT), nil); err != nil {
			logger.GetInstance().Error(fmt.Sprintf("Error starting Socket.IO server: %s", err))
			panic("Error starting Socket.IO server")
		}
	}()
	logger.GetInstance().Info("================================================================")
	logger.GetInstance().Info("=====================Application Started========================")
	logger.GetInstance().Info("================================================================")

	select {}
}

func CleanupUnfinishedTasks() {
	logger.GetInstance().Info("================================================================")
	logger.GetInstance().Info("=======================Application Stop=========================")
	logger.GetInstance().Info("================================================================")
	di.CleanupContainer()
	logger.CleanupQueuedLogs()
}
