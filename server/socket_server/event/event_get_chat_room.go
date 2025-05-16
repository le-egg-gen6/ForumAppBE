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

func RegisterEventGetChatRoom(router *socket_server.EventRouter, middleware ...socket_server.EventMiddlewareFunc) {
	router.RegisterEventHandler(constant.CSGetChatRoom, EventGetChatRoom, middleware...)
}

func EventGetChatRoom(client *socket_server.SocketClient, data *shared.SocketMessage) error {
	csGetChatRoom := utils.ConvertMessage[cs.CSGetChatRoom](data)
	if csGetChatRoom == nil {
		SendGetChatRoomFailure(client, "Bad request")
		return nil
	}
	user, err := repository.GetUserRepositoryInstance().FindByIDWithPreloadedField(client.UserID, "RoomChats")
	if err != nil {
		SendGetChatRoomFailure(client, "Unexpected error occurred, please try again")
		return nil
	}
	if user == nil {
		SendGetChatRoomFailure(client, "User not found")
		return nil
	}
	roomChatIds := make([]uint, 0)
	for _, roomChat := range user.RoomChats {
		roomChatIds = append(roomChatIds, roomChat.ID)
	}
	room := make([]*models.RoomChat, 0)
	for _, roomChatID := range roomChatIds {
		roomChat, err := repository.GetRoomChatRepositoryInstance().FindByIDWithPreloadedField(roomChatID, "Users")
		if err != nil {
			SendGetChatRoomFailure(client, "Unexpected error occurred, please try again")
			return nil
		}
		if roomChat != nil {
			room = append(room, roomChat)
		}
	}
	SendGetChatRoomSuccess(client, room)
	return nil
}

func SendGetChatRoomFailure(client *socket_server.SocketClient, message string) {
	utils.Send(client, constant.SCGetChatRoom, sc.SCGetChatRoom{Status: sc.StatusError, Message: message})
}

func SendGetChatRoomSuccess(client *socket_server.SocketClient, rooms []*models.RoomChat) {
	roomDtos := make([]dtos.RoomInfo, 0)
	for _, room := range rooms {
		roomDtos = append(roomDtos, *utils.ConvertToRoomInfo(room))
	}
	utils.Send(client, constant.SCGetChatRoom, sc.SCGetChatRoom{Status: sc.StatusSuccess, Rooms: roomDtos})
}
