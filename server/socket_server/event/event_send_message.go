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
		SendNewMessageFailure(client)
		return nil
	}
	sender, err := repository.GetUserRepositoryInstance().FindByID(client.UserID)
	if err != nil || sender == nil {
		SendNewMessageFailure(client)
		return nil
	}
	roomChat, err := repository.GetRoomChatRepositoryInstance().FindByIDWithPreloadedField(csNewMessage.RoomID, "Users", "Messages")
	if err != nil || roomChat == nil {
		SendNewMessageFailure(client)
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
		SendNewMessageFailure(client)
		return nil
	}
	roomMessage := &models.RoomMessage{
		UserID: sender.ID,
		RoomID: &roomChat.ID,
		Type:   constant.MessageTypeText,
		Body:   csNewMessage.Body,
	}
	roomMessage, err = repository.GetRoomMessageRepositoryInstance().Create(roomMessage)
	if err != nil {
		SendNewMessageFailure(client)
		return nil
	}
	roomChat.Messages = append(roomChat.Messages, roomMessage)
	err = repository.GetRoomChatRepositoryInstance().Update(roomChat)
	if err != nil {
		SendNewMessageFailure(client)
		return nil
	}
	BroadcastNewMessage(client, sender, roomChat, roomMessage)
	return nil
}

func SendNewMessageFailure(client *socket_server.SocketClient) {
	utils.Send(client, constant.SCNewMessage, sc.SCNewMessage{Status: sc.SendNewMessageFailed})
}

func BroadcastNewMessage(
	client *socket_server.SocketClient,
	sender *models.User,
	room *models.RoomChat,
	message *models.RoomMessage,
) {
	userIDs := make([]uint64, 0)
	for _, user := range room.Users {
		userIDs = append(userIDs, user.ID)
	}
	roomMemberSocketConn := client.Hub.GetClientsByUserIDs(userIDs)
	scMessage := sc.SCNewMessage{
		RoomID:      room.ID,
		Status:      sc.SendNewMessageSuccess,
		MessageInfo: *utils.ConvertToRoomMessageInfo(message, sender),
	}
	utils.Broadcast(roomMemberSocketConn, constant.SCNewMessage, scMessage)
}
