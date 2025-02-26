package service

import (
	"myproject/forum/dtos"
	"myproject/forum/models"
	repository2 "myproject/forum/repository"
)

type ICommentService interface {
	CreateComment(commentDTO *dtos.SimpleCommentDTO) (*dtos.CommentDTO, error)
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

func (cs *CommentService) CreateComment(commentDTO *dtos.SimpleCommentDTO) (*dtos.CommentDTO, error) {
	comment := &models.Comment{}
	comment.PostID = commentDTO.PostID
	comment.UserID = commentDTO.Author.ID
	comment.Body = commentDTO.Body

	createdComment, err := cs.CommentRepository.CreateComment(comment)
	if err != nil {
		return nil, err
	}

	response := dtos.CommentDTO{
		ID:        createdComment.ID,
		PostID:    createdComment.PostID,
		Author:    commentDTO.Author,
		Body:      createdComment.Body,
		CreatedAt: createdComment.CreatedAt,
		Reactions: []dtos.ReactionDTO{},
	}
	return &response, err

}
