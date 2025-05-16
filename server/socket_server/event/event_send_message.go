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
)

func RegisterEventNewMessage(router *socket_server.EventRouter, middleware ...socket_server.EventMiddlewareFunc) {
	router.RegisterEventHandler(constant.CSNewMessage, EventNewMessage, middleware...)
}

func EventNewMessage(client *socket_server.SocketClient, data *shared.SocketMessage) error {
	csNewMessage := utils.ConvertMessage[cs.CSNewMessage](data)
	if csNewMessage == nil {
		SendNewMessageFailure(client, "Bad request")
		return nil
	}
	sender, err := repository.GetUserRepositoryInstance().FindByID(client.UserID)
	if err != nil || sender == nil {
		SendNewMessageFailure(client, "User not found")
		return nil
	}
	roomChat, err := repository.GetRoomChatRepositoryInstance().FindByIDWithPreloadedField(csNewMessage.RoomID, "Users")
	if err != nil {
		SendNewMessageFailure(client, "Unexpected error occurred, please try again")
		return nil
	}
	if roomChat == nil {
		SendNewMessageFailure(client, "Room chat not found")
		return nil
	}
	inRoom := false
	for _, participant := range roomChat.Users {
		if participant.ID == sender.ID {
			inRoom = true
			break
		}
	}
	if !inRoom {
		SendNewMessageFailure(client, "Not in room")
		return nil
	}
	roomMessage := &models.RoomMessage{
		UserID: &sender.ID,
		RoomID: &roomChat.ID,
		Type:   constant.MessageTypeText,
		Body:   csNewMessage.Body,
	}
	roomMessage, err = repository.GetRoomMessageRepositoryInstance().Create(roomMessage)
	if err != nil {
		SendNewMessageFailure(client, "Unexpected error occurred, please try again")
		return nil
	}
	BroadcastRoomUpdateMessage(client, roomChat, roomMessage, sender)
	return nil
}

func SendNewMessageFailure(client *socket_server.SocketClient, message string) {
	utils.Send(client, constant.SCNewMessage, sc.SCUpdateMessage{Status: sc.StatusError, Message: message})
}
