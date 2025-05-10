package event

import (
	"forum/constant"
	"forum/models"
	"forum/repository"
	"forum/server/socket_server"
	"forum/server/socket_server/message/cs"
	"forum/server/socket_server/message/sc"
	"forum/shared"
	"forum/utils"
	"slices"
)

func RegisterEventCreateRoom(router *socket_server.EventRouter, middleware ...socket_server.EventMiddlewareFunc) {
	router.RegisterEventHandler(constant.CSCreateRoom, EventCreateRoom, middleware...)
}

func EventCreateRoom(client *socket_server.SocketClient, data *shared.SocketMessage) error {
	csCreateRoom := utils.ConvertMessage[cs.CSCreateRoom](data)
	if csCreateRoom == nil {
		SendCreateRoomFailure(client, "Bad request")
		return nil
	}
	if len(csCreateRoom.ParticipantIDs) < 2 {
		SendCreateRoomFailure(client, "Room member size must be greater than 2")
		return nil
	}
	if !slices.Contains(csCreateRoom.ParticipantIDs, client.UserID) {
		SendCreateRoomFailure(client, "You need in room to create room")
		return nil
	}
	users := make([]*models.User, 0)
	for _, userID := range csCreateRoom.ParticipantIDs {
		user, err := repository.GetUserRepositoryInstance().FindByIDWithPreloadedField(userID, "RoomChats")
		if err != nil || user == nil {
			SendCreateRoomFailure(client, "Not in room")
			return nil
		}
		users = append(users, user)
	}
	roomChat := &models.RoomChat{
		Name: csCreateRoom.Name,
		Type: constant.RoomTypeGroup,
	}
	if len(csCreateRoom.ParticipantIDs) == 2 {
		roomChat.Type = constant.RoomTypePrivate
	}
	roomChat, err := repository.GetRoomChatRepositoryInstance().Create(roomChat)
	if err != nil {
		SendCreateRoomFailure(client, "Room not found")
		return nil
	}
	err = repository.GetRoomChatRepositoryInstance().UpdateAssociations(roomChat, "Users", users)
	if err != nil {
		SendCreateRoomFailure(client, "Unexpected error occurred, please try again")
		return nil
	}
	SendCreateRoomSuccess(client, roomChat)
	BroadcastUpdateRoomInfo(client, roomChat)
	return nil
}

func SendCreateRoomFailure(client *socket_server.SocketClient, message string) {
	utils.Send(client, constant.SCCreateRoom, sc.SCCreateRoom{Status: sc.StatusError, Message: message})
}

func SendCreateRoomSuccess(client *socket_server.SocketClient, room *models.RoomChat) {
	utils.Send(client, constant.SCCreateRoom, sc.SCCreateRoom{Status: sc.StatusSuccess, RoomInfo: *utils.ConvertToRoomInfo(room)})
}
