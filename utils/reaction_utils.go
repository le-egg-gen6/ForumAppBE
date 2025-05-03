package utils

import (
	"forum/constant"
	"forum/dtos"
	"forum/models"
)

func IsReactionTypeValid(reactionType string) bool {
	valid, ok := constant.AllowedReactionTypes[reactionType]
	return valid && ok
}

func ConvertToReactionDTO(reaction *models.ContentReaction) *dtos.ReactionDTO {
	return &dtos.ReactionDTO{
		Type:  reaction.Type,
		Count: reaction.Count,
	}
}
