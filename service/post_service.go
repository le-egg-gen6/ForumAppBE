package service

import (
	"myproject/forum/dtos"
	repository2 "myproject/forum/repository"
)

type IPostService interface {
	CreatePost(postDTO *dtos.SimplePostDTO) (*dtos.PostDTO, error)
}

type PostService struct {
	UserRepository     repository2.IUserRepository
	PostRepository     repository2.IPostRepository
	CommentRepository  repository2.ICommentRepository
	ReactionRepository repository2.IReactionRepository
}

func NewPostService(
	userRepository repository2.IUserRepository,
	postRepository repository2.IPostRepository,
	commentRepository repository2.ICommentRepository,
	reactionRepository repository2.IReactionRepository,
) *PostService {
	return &PostService{
		UserRepository:     userRepository,
		PostRepository:     postRepository,
		CommentRepository:  commentRepository,
		ReactionRepository: reactionRepository,
	}
}

func (ps *PostService) CreatePost(postDTO *dtos.SimplePostDTO) (*dtos.PostDTO, error) {

}
