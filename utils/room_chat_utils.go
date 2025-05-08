package utils

import (
	"forum/dtos"
	"forum/models"
)

func ConvertToRoomInfo(room *models.RoomChat) *dtos.RoomInfo {
	roomInfo := &dtos.RoomInfo{
		RoomID: room.ID,
		Name:   room.Name,
		Type:   room.Type,
	}
	participants := make([]dtos.SimpleUserDTO, 0)
	for _, user := range room.Users {
		participants = append(participants, *ConvertToSimpleUserDTO(user))
	}
	roomInfo.ParticipantInfos = participants
	return roomInfo
}
