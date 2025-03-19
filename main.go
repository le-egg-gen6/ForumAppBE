package main

import (
	"fmt"
	"log"
	"myproject/forum/config"
	"myproject/forum/di"
	"myproject/forum/logger"
	"myproject/forum/router"
	"myproject/forum/tcp"
	"net/http"
)

func main() {
	logger.InitializeLogger()
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
		log.Println("HTTP Server running on port", cfg.PORT)
		if err := initializeRouter.Run(fmt.Sprintf(":%d", cfg.PORT)); err != nil {
			log.Fatalf("Error starting HTTP server: %s", err)
		}
	}()

	socketServer := tcp.NewSocketServer()
	go func() {
		http.Handle("/socket.io/", socketServer)
		log.Println("Socket.IO server running on port", cfg.TCP_PORT)
		if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.TCP_PORT), nil); err != nil {
			log.Fatalf("Error starting Socket.IO server: %s", err)
		}
	}()

	select {}
}

func CleanupUnfinishedTasks() {
	di.CleanupContainer()
	logger.CleanupQueuedLogs()
}
