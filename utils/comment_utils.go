package utils

import (
	"forum/dtos"
	"forum/models"
)

func ConvertToCommentDTO(user *models.User, comment *models.Comment) *dtos.CommentDTO {
	commentDTO := &dtos.CommentDTO{
		ID:        comment.ID,
		PostID:    *comment.PostID,
		Author:    *ConvertToSimpleUserDTO(user),
		Body:      comment.Body,
		CreatedAt: *comment.CreatedAt,
	}
	reactions := make([]dtos.ReactionDTO, 0)
	for _, reaction := range comment.Reactions {
		reactions = append(reactions, *ConvertToReactionDTO(reaction))
	}
	commentDTO.Reactions = reactions
	return commentDTO
}

func ConvertToSimpleCommentDTO(user *models.User, comment *models.Comment) *dtos.SimpleCommentDTO {
	return &dtos.SimpleCommentDTO{
		ID:        comment.ID,
		PostID:    *comment.PostID,
		Author:    *ConvertToSimpleUserDTO(user),
		Body:      comment.Body,
		CreatedAt: *comment.CreatedAt,
	}
}
