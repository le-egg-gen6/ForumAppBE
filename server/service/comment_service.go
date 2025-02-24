package service

import "myproject/forum/server/repository"

type ICommentService interface {
}

type CommentService struct {
	CommentRepository  repository.ICommentRepository
	ReactionRepository repository.IReactionRepository
}

func NewCommentService(commentRepository repository.ICommentRepository, reactionRepository repository.IReactionRepository) *CommentService {
	return &CommentService{
		CommentRepository:  commentRepository,
		ReactionRepository: reactionRepository,
	}
}
