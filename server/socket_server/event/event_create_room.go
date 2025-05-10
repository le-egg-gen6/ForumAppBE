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
		SendCreateRoomFailure(client)
		return nil
	}
	if len(csCreateRoom.ParticipantIDs) < 2 || !slices.Contains(csCreateRoom.ParticipantIDs, client.UserID) {
		SendCreateRoomFailure(client)
		return nil
	}
	users := make([]*models.User, 0)
	for _, userID := range csCreateRoom.ParticipantIDs {
		user, err := repository.GetUserRepositoryInstance().FindByIDWithPreloadedField(userID, "RoomChats")
		if err != nil || user == nil {
			SendCreateRoomFailure(client)
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
		SendCreateRoomFailure(client)
		return nil
	}
	err = repository.GetRoomChatRepositoryInstance().UpdateAssociations(roomChat, "Users", users)
	if err != nil {
		SendCreateRoomFailure(client)
		return nil
	}
	BroadcastCreateRoomSuccess(client, csCreateRoom.ParticipantIDs, roomChat)
	return nil
}

func SendCreateRoomFailure(client *socket_server.SocketClient) {
	utils.Send(client, constant.SCCreateRoom, sc.SCCreateRoom{Status: sc.RoomCreateFailed})
}

func BroadcastCreateRoomSuccess(
	client *socket_server.SocketClient,
	userIDs []uint,
	room *models.RoomChat,
) {
	roomMembersSocketConn := client.Hub.GetClientsByUserIDs(userIDs)
	scMessage := sc.SCCreateRoom{Status: sc.RoomCreatedSuccess, RoomInfo: *utils.ConvertToRoomInfo(room)}
	utils.Broadcast(roomMembersSocketConn, constant.SCCreateRoom, scMessage)
}
