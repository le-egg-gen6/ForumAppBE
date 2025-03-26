package service

import (
	"forum/dtos"
	"forum/models"
	repository2 "forum/repository"
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
	post := &models.Post{}
	post.AuthorID = postDTO.Author.ID
	post.Content = postDTO.Content

	createdPost, err := ps.PostRepository.CreatePost(post)
	if err != nil {
		return nil, err
	}

	response := dtos.PostDTO{
		ID:          createdPost.ID,
		Content:     createdPost.Content,
		Author:      postDTO.Author,
		CreatedAt:   createdPost.CreatedAt,
		Reactions:   []dtos.ReactionDTO{}, // Assuming you might populate this later
		TopComments: []dtos.CommentDTO{},  // Assuming you might populate this later
	}

	return &response, err
}
