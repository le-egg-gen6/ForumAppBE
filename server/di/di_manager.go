package di

import (
	"gorm.io/gorm"
	"myproject/forum/server/config"
	"myproject/forum/server/controller"
	"myproject/forum/server/repository"
	"myproject/forum/server/router/routes"
	"myproject/forum/server/service"
)

type Container struct {
	DB *gorm.DB

	UserRepository *repository.UserRepository
	UserService    *service.UserService
	UserController *controller.UserController
	UserRoutes     *routes.UserRoutes

	PostRepository *repository.PostRepository
	PostService    *service.PostService
	PostController *controller.PostController
	PostRoutes     *routes.PostRoutes
}

func InitializeContainer(cfg *config.Config) *Container {
	db := config.ConnectDB(cfg)

	//User
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)
	userRoutes := &routes.UserRoutes{
		UserController: userController,
	}

	postRepository := repository.NewPostRepository(db)
	postService := service.NewPostService(postRepository)
	postController := controller.NewPostController(postService)
	postRoutes := &routes.PostRoutes{
		PostController: postController,
	}

	return &Container{
		DB:             db,
		UserRepository: userRepository,
		UserService:    userService,
		UserController: userController,
		UserRoutes:     userRoutes,
		PostRepository: postRepository,
		PostService:    postService,
		PostController: postController,
		PostRoutes:     postRoutes,
	}
}

func CleanupContainer() {
}
