package service

import "myproject/forum/server/repository"

type IPostService interface {
}

type PostService struct {
	PostRepository repository.IPostRepository
}

func NewPostService(postRepository repository.IPostRepository) *PostService {
	return &PostService{
		PostRepository: postRepository,
	}
}
