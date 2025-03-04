package main

import (
	"fmt"
	"log"
	"myproject/forum/config"
	"myproject/forum/di"
	"myproject/forum/logger"
	"myproject/forum/router"
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

	if err := initializeRouter.Run(fmt.Sprintf(":%d", cfg.PORT)); err != nil {
		log.Fatalf("Error starting server, %s", err)
	}

}

func CleanupUnfinishedTasks() {
	di.CleanupContainer()
	logger.CleanupQueuedLogs()
}
