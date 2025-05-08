package utils

import (
	"forum/dtos"
	"forum/models"
)

func ConvertToRoomMessageInfo(message *models.RoomMessage, user *models.User) *dtos.MessageInfo {
	messageInfo := &dtos.MessageInfo{
		MessageID: message.ID,
		Body:      message.Body,
		Type:      message.Type,
		CreatedAt: *message.CreatedAt,
	}
	if user != nil {
		messageInfo.Author = *ConvertToSimpleUserDTO(user)
	}
	reactionDTOs := make([]dtos.ReactionDTO, 0)
	for _, reaction := range message.Reactions {
		reactionDTOs = append(reactionDTOs, *ConvertToReactionDTO(reaction))
	}
	messageInfo.Reactions = reactionDTOs
	return messageInfo
}
