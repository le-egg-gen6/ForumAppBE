package service

import (
	repository2 "myproject/forum/repository"
)

type IPostService interface {
}

type PostService struct {
	PostRepository     repository2.IPostRepository
	ReactionRepository repository2.IReactionRepository
}

func NewPostService(postRepository repository2.IPostRepository, reactionRepository repository2.IReactionRepository) *PostService {
	return &PostService{
		PostRepository:     postRepository,
		ReactionRepository: reactionRepository,
	}
}
