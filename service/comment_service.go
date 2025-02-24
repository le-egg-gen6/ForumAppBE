package service

import (
	repository2 "myproject/forum/repository"
)

type ICommentService interface {
}

type CommentService struct {
	CommentRepository  repository2.ICommentRepository
	ReactionRepository repository2.IReactionRepository
}

func NewCommentService(commentRepository repository2.ICommentRepository, reactionRepository repository2.IReactionRepository) *CommentService {
	return &CommentService{
		CommentRepository:  commentRepository,
		ReactionRepository: reactionRepository,
	}
}
