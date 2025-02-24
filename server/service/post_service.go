package service

import "myproject/forum/server/repository"

type IPostService interface {
}

type PostService struct {
	PostRepository     repository.IPostRepository
	ReactionRepository repository.IReactionRepository
}

func NewPostService(postRepository repository.IPostRepository, reactionRepository repository.IReactionRepository) *PostService {
	return &PostService{
		PostRepository:     postRepository,
		ReactionRepository: reactionRepository,
	}
}
