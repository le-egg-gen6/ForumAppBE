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

func RegisterEventLeaveRoom(router *socket_server.EventRouter, middleware ...socket_server.EventMiddlewareFunc) {
	router.RegisterEventHandler(constant.CSLeaveRoom, EventLeaveRoom, middleware...)
}

func EventLeaveRoom(client *socket_server.SocketClient, data *shared.SocketMessage) error {
	csLeaveRoom := utils.ConvertMessage[cs.CSLeaveRoom](data)
	if csLeaveRoom == nil {
		SendLeaveRoomFailure(client)
		return nil
	}
	user, err := repository.GetUserRepositoryInstance().FindByIDWithPreloadedField(client.UserID, "RoomChats")
	if err != nil || user == nil {
		SendLeaveRoomFailure(client)
		return nil
	}
	roomChat, err := repository.GetRoomChatRepositoryInstance().FindByIDWithPreloadedField(csLeaveRoom.RoomID, "Users")
	if err != nil || roomChat == nil {
		SendLeaveRoomFailure(client)
		return nil
	}
	if roomChat.Type == constant.RoomTypePrivate {
		SendLeaveRoomFailure(client)
		return nil
	}

}

func SendLeaveRoomFailure(client *socket_server.SocketClient) {
	utils.Send(client, constant.SCLeaveRoom, sc.SCLeaveRoom{Status: sc.RoomLeaveFailed})
}

func SendLeaveRoomSuccess(client *socket_server.SocketClient) {
	utils.Send(client, constant.SCLeaveRoom, sc.SCLeaveRoom{Status: sc.RoomLeaveSuccess})
}

func BroadcastLeaveRoomSuccess(
	client *socket_server.SocketClient,
	userIDs []int,
	message *models.RoomMessage,
) {

}
