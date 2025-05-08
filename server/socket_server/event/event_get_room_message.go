package event

import (
	"forum/constant"
	"forum/dtos"
	"forum/repository"
	"forum/server/socket_server"
	"forum/server/socket_server/message/cs"
	"forum/server/socket_server/message/sc"
	"forum/shared"
	"forum/utils"
)

func RegisterEventGetRoomMessage(router *socket_server.EventRouter, middlewares ...socket_server.EventMiddlewareFunc) {
	router.RegisterEventHandler(constant.CSGetRoomMessage, EventGetRoomMessage, middlewares...)
}

func EventGetRoomMessage(client *socket_server.SocketClient, data *shared.SocketMessage) error {
	csGetRoomMessage := utils.ConvertMessage[cs.CSGetRoomMessage](data)
	messages := make([]dtos.MessageInfo, 0)
	if csGetRoomMessage == nil {
		SendGetRoomMessage(client, messages)
		return nil
	}
	roomChat, err := repository.GetRoomChatRepositoryInstance().FindByIDWithPreloadedField(csGetRoomMessage.RoomID, "Messages", "Users")
	if err != nil || roomChat == nil {
		SendGetRoomMessage(client, messages)
		return nil
	}
	inRoom := false
	for _, user := range roomChat.Users {
		if user.ID == client.UserID {
			inRoom = true
			break
		}
	}
	if !inRoom {
		SendGetRoomMessage(client, messages)
		return nil
	}
	for _, message := range roomChat.Messages {
		if message.Type == constant.MessageTypeText {
			user, err := repository.GetUserRepositoryInstance().FindByID(message.UserID)
			if err != nil || user == nil {
				continue
			}
			messages = append(messages, *utils.ConvertToRoomMessageInfo(message, user))
		} else {
			messages = append(messages, *utils.ConvertToRoomMessageInfo(message, nil))
		}
	}
	SendGetRoomMessage(client, messages)
	return nil
}

func SendGetRoomMessage(client *socket_server.SocketClient, messages []dtos.MessageInfo) {
	utils.Send(client, constant.SCGetRoomMessage, sc.SCGetRoomMessage{Messages: messages})
}
