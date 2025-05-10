package event

import (
	"fmt"
	"forum/constant"
	"forum/models"
	"forum/repository"
	"forum/server/socket_server"
	"forum/server/socket_server/message/cs"
	"forum/server/socket_server/message/sc"
	"forum/shared"
	"forum/utils"
)

func RegisterEventAddParticipantRoomChat(router *socket_server.EventRouter, middlewares ...socket_server.EventMiddlewareFunc) {
	router.RegisterEventHandler(constant.CSAddParticipantRoomChat, EventAddParticipantRoomChat, middlewares...)
}

func EventAddParticipantRoomChat(client *socket_server.SocketClient, data *shared.SocketMessage) error {
	csAddParticipantRoomChat := utils.ConvertMessage[cs.CSAddParticipantRoomChat](data)
	if csAddParticipantRoomChat == nil {
		SendAddParticipantRoomChatFailure(client, "Bad request")
		return nil
	}
	if client.UserID == csAddParticipantRoomChat.ParticipantID {
		SendCreateRoomFailure(client, "Can't add yourself")
		return nil
	}
	roomChat, err := repository.GetRoomChatRepositoryInstance().FindByIDWithPreloadedField(csAddParticipantRoomChat.RoomID, "Users")
	if err != nil {
		SendAddParticipantRoomChatFailure(client, "Unexpected error occurred, please try again")
		return nil
	}
	if roomChat == nil {
		SendAddParticipantRoomChatFailure(client, "Room not exist")
		return nil
	}
	inRoom := false
	for _, user_ := range roomChat.Users {
		if user_.ID == csAddParticipantRoomChat.ParticipantID {
			inRoom = true
			break
		}
	}
	if inRoom {
		SendAddParticipantRoomChatFailure(client, "Can't add yourself")
		return nil
	}
	user, err := repository.GetUserRepositoryInstance().FindByID(client.UserID)
	if err != nil {
		SendAddParticipantRoomChatFailure(client, "Unexpected error occurred, please try again")
		return nil
	}
	if user == nil {
		SendAddParticipantRoomChatFailure(client, "User not exist")
		return nil
	}
	err = repository.GetRoomChatRepositoryInstance().UpdateAssociations(roomChat, "Users", user)
	if err != nil {
		SendAddParticipantRoomChatFailure(client, "Unexpected error occurred, please try again")
		return nil
	}
	SendAddParticipantRoomChatSuccess(client)
	roomMessage := &models.RoomMessage{
		RoomID: &roomChat.ID,
		Type:   constant.MessageTypeNotification,
		Body:   fmt.Sprintf("%s joined chat room.", user.Username),
	}
	roomMessage, err = repository.GetRoomMessageRepositoryInstance().Create(roomMessage)
	if err != nil || roomMessage == nil {
		return nil
	}
	BroadcastRoomNewMessage(client, roomChat, roomMessage, nil)
	roomChat.Users = append(roomChat.Users, user)
	BroadcastUpdateRoomInfo(client, roomChat)
	return nil
}

func SendAddParticipantRoomChatFailure(client *socket_server.SocketClient, message string) {
	utils.Send(client, constant.SCAddParticipantRoomChat, sc.SCAddParticipantRoomChat{Status: sc.StatusError, Message: message})
}

func SendAddParticipantRoomChatSuccess(client *socket_server.SocketClient) {
	utils.Send(client, constant.SCAddParticipantRoomChat, sc.SCAddParticipantRoomChat{Status: sc.StatusSuccess})
}
