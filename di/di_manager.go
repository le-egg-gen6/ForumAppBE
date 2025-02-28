package di

import (
	"gorm.io/gorm"
	"myproject/forum/cache/redis"
	"myproject/forum/cloudinary"
	config2 "myproject/forum/config"
	controller2 "myproject/forum/controller"
	"myproject/forum/mail_sender"
	repository2 "myproject/forum/repository"
	routes2 "myproject/forum/router/routes"
	service2 "myproject/forum/service"
)

type Container struct {
	DB *gorm.DB

	FileUploader *cloudinary.FileUploader
	MailSender   *mail_sender.MailSender
	RedisClient  *redis.RedisClient

	UserRepository     *repository2.UserRepository
	PostRepository     *repository2.PostRepository
	CommentRepository  *repository2.CommentRepository
	ReactionRepository *repository2.ReactionRepository

	UserService     *service2.UserService
	PostService     *service2.PostService
	CommentService  *service2.CommentService
	ReactionService *service2.ReactionService

	AuthController     *controller2.AuthController
	UserController     *controller2.UserController
	PostController     *controller2.PostController
	CommentController  *controller2.CommentController
	ReactionController *controller2.ReactionController

	AuthRoutes     *routes2.AuthRoutes
	UserRoutes     *routes2.UserRoutes
	PostRoutes     *routes2.PostRoutes
	CommentRoutes  *routes2.CommentRoutes
	ReactionRoutes *routes2.ReactionRoutes
}

func InitializeContainer(cfg *config2.Config) *Container {
	db := config2.ConnectDB(cfg)

	//File uploader
	fileUploader := cloudinary.NewFileUploader()

	//Mail Sender
	mailSender := mail_sender.NewMailSender()

	//Redis Client
	redisClient := redis.NewRedisClient(redis.LoadRedisConfig(cfg))

	//Repository
	userRepository := repository2.NewUserRepository(db)
	postRepository := repository2.NewPostRepository(db)
	commentRepository := repository2.NewCommentRepository(db)
	reactionRepository := repository2.NewReactionRepository(db)

	//Service
	userService := service2.NewUserService(userRepository)
	postService := service2.NewPostService(
		userRepository,
		postRepository,
		commentRepository,
		reactionRepository,
	)
	commentService := service2.NewCommentService(
		userRepository,
		commentRepository,
		reactionRepository,
	)
	reactionService := service2.NewReactionService(reactionRepository)

	//Controller
	authController := controller2.NewAuthController(userService)
	userController := controller2.NewUserController(userService)
	postController := controller2.NewPostController(postService)
	commentController := controller2.NewCommentController(commentService)
	reactionController := controller2.NewReactionController(reactionService)

	//Routes
	authRoutes := &routes2.AuthRoutes{
		AuthController: authController,
	}
	userRoutes := &routes2.UserRoutes{
		UserController: userController,
	}
	postRoutes := &routes2.PostRoutes{
		PostController: postController,
	}
	reactionRoutes := &routes2.ReactionRoutes{
		ReactionController: reactionController,
	}
	commentRoutes := &routes2.CommentRoutes{
		CommentController: commentController,
	}

	return &Container{
		DB: db,

		FileUploader: fileUploader,
		MailSender:   mailSender,
		RedisClient:  redisClient,

		ReactionRepository: reactionRepository,
		UserRepository:     userRepository,
		PostRepository:     postRepository,
		CommentRepository:  commentRepository,

		UserService:     userService,
		PostService:     postService,
		CommentService:  commentService,
		ReactionService: reactionService,

		AuthController:     authController,
		UserController:     userController,
		PostController:     postController,
		CommentController:  commentController,
		ReactionController: reactionController,

		AuthRoutes:     authRoutes,
		UserRoutes:     userRoutes,
		PostRoutes:     postRoutes,
		CommentRoutes:  commentRoutes,
		ReactionRoutes: reactionRoutes,
	}
}

func CleanupContainer() {
}
