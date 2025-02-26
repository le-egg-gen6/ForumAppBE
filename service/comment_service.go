package service

import (
	repository2 "myproject/forum/repository"
)

type ICommentService interface {
}

type CommentService struct {
	UserRepository     repository2.IUserRepository
	CommentRepository  repository2.ICommentRepository
	ReactionRepository repository2.IReactionRepository
}

func NewCommentService(
	userRepository repository2.IUserRepository,
	commentRepository repository2.ICommentRepository,
	reactionRepository repository2.IReactionRepository,
) *CommentService {
	return &CommentService{
		UserRepository:     userRepository,
		CommentRepository:  commentRepository,
		ReactionRepository: reactionRepository,
	}
}
