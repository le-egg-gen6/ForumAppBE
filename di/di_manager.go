package di

import (
	"gorm.io/gorm"
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

	MailSender *mail_sender.MailSender

	UserRepository *repository2.UserRepository
	UserService    *service2.UserService
	UserController *controller2.UserController
	UserRoutes     *routes2.UserRoutes

	PostRepository *repository2.PostRepository
	PostService    *service2.PostService
	PostController *controller2.PostController
	PostRoutes     *routes2.PostRoutes

	CommentRepository *repository2.CommentRepository
	CommentService    *service2.CommentService
	CommentController *controller2.CommentController
	CommentRoutes     *routes2.CommentRoutes

	ReactionRepository *repository2.ReactionRepository
	ReactionService    *service2.ReactionService
	ReactionController *controller2.ReactionController
	ReactionRoutes     *routes2.ReactionRoutes
}

func InitializeContainer(cfg *config2.Config) *Container {
	db := config2.ConnectDB(cfg)

	//File uploader
	fileUploader := cloudinary.NewFileUploader()

	//Mail Sender
	mailSender := mail_sender.NewMailSender()

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
	userController := controller2.NewUserController(userService)
	postController := controller2.NewPostController(postService)
	commentController := controller2.NewCommentController(commentService)
	reactionController := controller2.NewReactionController(reactionService)

	//Routes
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

		MailSender: mailSender,

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
