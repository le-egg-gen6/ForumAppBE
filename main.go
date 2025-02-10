package ForumApp

import (
	"log"
	"myproject/forum/server/config"
	"myproject/forum/server/di"
	"myproject/forum/server/logger"
	"myproject/forum/server/router"
)

func main() {
	logger.InitializeLogger()
	defer CleanupUnfinishedTasks()

	cfg := config.LoadConfig()

	container := di.InitializeContainer(cfg)

	routesModules := []router.Router{
		container.UserRoutes,
		container.PostRoutes,
	}

	initializeRouter := router.InitializeRouter(cfg, routesModules)

	if err := initializeRouter.Run(":" + cfg.PORT); err != nil {
		log.Fatalf("Error starting server, %s", err)
	}

}

func CleanupUnfinishedTasks() {
	di.CleanupContainer()
	logger.CleanupQueuedLogs()
}
