package ForumApp

import (
	"log"
	"myproject/forum/server/config"
	"myproject/forum/server/di"
	"myproject/forum/server/router"
)

func main() {
	cfg := config.LoadConfig()

	container := di.InitializeContainer(cfg)

	routesModules := []router.Router{
		container.UserRoutes,
	}

	initializeRouter := router.InitializeRouter(cfg, routesModules)

	if err := initializeRouter.Run(":" + cfg.PORT); err != nil {
		log.Fatalf("Error starting server, %s", err)
	}

}
