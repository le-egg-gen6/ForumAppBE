package utils

import (
	"forum/dtos"
	"forum/models"
)

func ConvertToPostDTO(
	user *models.User,
	post *models.Post,
	topComments []*models.Comment,
	topCommentAuthors []*models.User,

) *dtos.PostDTO {
	postDTO := &dtos.PostDTO{
		ID:        post.ID,
		Content:   post.Content,
		Author:    *ConvertToSimpleUserDTO(user),
		CreatedAt: post.CreatedAt,
	}

	images := make([]string, 0)
	for _, image := range post.Images {
		images = append(images, image.URL)
	}
	postDTO.Images = images

	reactionDTOs := make([]dtos.ReactionDTO, 0)
	for _, reaction := range post.Reactions {
		reactionDTOs = append(reactionDTOs, *ConvertToReactionDTO(reaction))
	}
	postDTO.Reactions = reactionDTOs

	topCommentDTOs := make([]dtos.CommentDTO, 0)
	for i, comment := range topComments {
		topCommentDTOs = append(topCommentDTOs, *ConvertToCommentDTO(topCommentAuthors[i], comment))
	}
	postDTO.TopComments = topCommentDTOs
	return postDTO
}

func ConvertToSimplePostDTO(user *models.User, post *models.Post) *dtos.SimplePostDTO {
	simplePostDTO := &dtos.SimplePostDTO{
		ID:        post.ID,
		Content:   post.Content,
		Author:    *ConvertToSimpleUserDTO(user),
		CreatedAt: post.CreatedAt,
	}
	images := make([]string, 0)
	for _, image := range post.Images {
		images = append(images, image.URL)
	}
	simplePostDTO.Images = images
	return simplePostDTO
}
