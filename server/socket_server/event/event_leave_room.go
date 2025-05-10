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

func RegisterEventLeaveRoom(router *socket_server.EventRouter, middleware ...socket_server.EventMiddlewareFunc) {
	router.RegisterEventHandler(constant.CSLeaveRoom, EventLeaveRoom, middleware...)
}

func EventLeaveRoom(client *socket_server.SocketClient, data *shared.SocketMessage) error {
	csLeaveRoom := utils.ConvertMessage[cs.CSLeaveRoom](data)
	if csLeaveRoom == nil {
		SendLeaveRoomFailure(client, "Bad request")
		return nil
	}
	user, err := repository.GetUserRepositoryInstance().FindByID(client.UserID)
	if err != nil {
		SendLeaveRoomFailure(client, "Unexpected error occurred, please try again")
		return nil
	}
	if user == nil {
		SendLeaveRoomFailure(client, "User not found")
		return nil
	}
	roomChat, err := repository.GetRoomChatRepositoryInstance().FindByIDWithPreloadedField(csLeaveRoom.RoomID, "Users")
	if err != nil {
		SendLeaveRoomFailure(client, "Unexpected error occurred, please try again")
		return nil
	}
	if roomChat == nil {
		SendLeaveRoomFailure(client, "Room chat not found")
		return nil
	}
	if roomChat.Type == constant.RoomTypePrivate {
		SendLeaveRoomFailure(client, "Can not leave private chat")
		return nil
	}
	inRoom := false
	for _, user_ := range roomChat.Users {
		if user_.ID == user.ID {
			inRoom = true
			break
		}
	}
	if !inRoom {
		SendLeaveRoomFailure(client, "Not in room")
		return nil
	}
	err = repository.GetRoomChatRepositoryInstance().DeleteAssociations(roomChat, "Users", user)
	if err != nil {
		SendLeaveRoomFailure(client, "Unexpected error occurred, please try again")
		return nil
	}
	SendLeaveRoomSuccess(client)
	users := make([]*models.User, 0)
	for _, user_ := range roomChat.Users {
		if user_.ID != user.ID {
			users = append(users, user_)
		}
	}
	roomChat.Users = users
	BroadcastUpdateRoomInfo(client, roomChat)
	roomMessage := &models.RoomMessage{
		RoomID: &roomChat.ID,
		Type:   constant.MessageTypeNotification,
		Body:   fmt.Sprintf("%s leaved chat room.", user.Username),
	}
	roomMessage, err = repository.GetRoomMessageRepositoryInstance().Create(roomMessage)
	if err != nil || roomMessage == nil {
		return nil
	}
	BroadcastRoomNewMessage(client, roomChat, roomMessage, nil)
	return nil
}

func SendLeaveRoomFailure(client *socket_server.SocketClient, message string) {
	utils.Send(client, constant.SCLeaveRoom, sc.SCLeaveRoom{Status: sc.StatusError, Message: message})
}

func SendLeaveRoomSuccess(client *socket_server.SocketClient) {
	utils.Send(client, constant.SCLeaveRoom, sc.SCLeaveRoom{Status: sc.StatusSuccess})
}
