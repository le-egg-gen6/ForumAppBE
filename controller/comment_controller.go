package controller

import (
	"myproject/forum/service"
)

type CommentController struct {
	CommentService *service.CommentService
}

func NewCommentController(commentService *service.CommentService) *CommentController {
	return &CommentController{
		CommentService: commentService,
	}
}
