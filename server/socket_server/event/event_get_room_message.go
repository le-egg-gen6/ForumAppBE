package event

import (
	"forum/constant"
	"forum/dtos"
	"forum/models"
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
		SendGetRoomMessageFailure(client, "Bad request")
		return nil
	}
	roomChat, err := repository.GetRoomChatRepositoryInstance().FindByIDWithPreloadedField(csGetRoomMessage.RoomID, "Messages", "Users")
	if err != nil {
		SendGetRoomMessageFailure(client, "Unexpected error occurred, please try again")
		return nil
	}
	if roomChat == nil {
		SendGetRoomMessageFailure(client, "Room chat not found")
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
		SendGetRoomMessageFailure(client, "Not in room")
		return nil
	}
	mapUserIdToUser := make(map[uint]*models.User)
	for _, message := range roomChat.Messages {
		if message.Type == constant.MessageTypeText {
			user_, ok := mapUserIdToUser[*message.UserID]
			if !ok {
				user_, err = repository.GetUserRepositoryInstance().FindByIDWithPreloadedField(*message.UserID, "Avatar")
				if err != nil {
					SendGetRoomMessageFailure(client, "Unexpected error occurred, please try again")
					return nil
				}
				mapUserIdToUser[*message.UserID] = user_
			}
			messages = append(messages, *utils.ConvertToRoomMessageInfo(message, user_))
		} else {
			messages = append(messages, *utils.ConvertToRoomMessageInfo(message, nil))
		}
	}
	SendGetRoomMessageSuccess(client, messages)
	return nil
}

func SendGetRoomMessageFailure(client *socket_server.SocketClient, message string) {
	utils.Send(client, constant.SCGetRoomMessage, sc.SCGetRoomMessage{Status: sc.StatusError, Message: message})
}
func SendGetRoomMessageSuccess(client *socket_server.SocketClient, messages []dtos.MessageInfo) {
	utils.Send(client, constant.SCGetRoomMessage, sc.SCGetRoomMessage{Status: sc.StatusSuccess, Messages: messages})
}
