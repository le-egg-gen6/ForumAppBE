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

	CommentRepository *repository.CommentRepository
	CommentService    *service.CommentService
	CommentController *controller.CommentController
	CommentRoutes     *routes.CommentRoutes

	ReactionRepository *repository.ReactionRepository
	ReactionService    *service.ReactionService
	ReactionController *controller.ReactionController
	ReactionRoutes     *routes.ReactionRoutes
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

	//Reaction
	reactionRepository := repository.NewReactionRepository(db)
	reactionService := service.NewReactionService(reactionRepository)
	reactionController := controller.NewReactionController(reactionService)
	reactionRoutes := &routes.ReactionRoutes{
		ReactionController: reactionController,
	}

	//Post
	postRepository := repository.NewPostRepository(db)
	postService := service.NewPostService(postRepository, reactionRepository)
	postController := controller.NewPostController(postService)
	postRoutes := &routes.PostRoutes{
		PostController: postController,
	}

	//Comment
	commentRepository := repository.NewCommentRepository(db)
	commentService := service.NewCommentService(commentRepository, reactionRepository)
	commentController := controller.NewCommentController(commentService)
	commentRoutes := &routes.CommentRoutes{
		CommentController: commentController,
	}

	return &Container{
		DB: db,

		ReactionRepository: reactionRepository,
		ReactionService:    reactionService,
		ReactionController: reactionController,
		ReactionRoutes:     reactionRoutes,

		UserRepository: userRepository,
		UserService:    userService,
		UserController: userController,
		UserRoutes:     userRoutes,

		PostRepository: postRepository,
		PostService:    postService,
		PostController: postController,
		PostRoutes:     postRoutes,

		CommentRepository: commentRepository,
		CommentService:    commentService,
		CommentController: commentController,
		CommentRoutes:     commentRoutes,
	}
}

func CleanupContainer() {
}
