package service

import "myproject/forum/server/repository"

type ICommentService interface {
}

type CommentService struct {
	CommentRepository *repository.CommentRepository
}

func NewCommentService(commentRepository *repository.CommentRepository) *CommentService {
	return &CommentService{
		CommentRepository: commentRepository,
	}
}
