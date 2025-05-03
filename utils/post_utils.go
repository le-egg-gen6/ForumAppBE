package utils

import (
	"forum/dtos"
	"forum/models"
)

func ConvertToPostDTO(post *models.Post) *dtos.PostDTO {
	return &dtos.PostDTO{}
}

func ConvertToSimplePostDTO(user *models.User, post *models.Post) *dtos.SimplePostDTO {
	return &dtos.SimplePostDTO{
		ID:      post.ID,
		Content: post.Content,
		Author:  ConvertToSimpleUserDTO(user),
	}
}
