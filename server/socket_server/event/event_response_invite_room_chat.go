package event

import (
	"forum/constant"
	"forum/server/socket_server"
	"forum/shared"
)

func RegisterEventResponseInviteRoomChat(router *socket_server.EventRouter, middlewares ...socket_server.EventMiddlewareFunc) {
	router.RegisterEventHandler(constant.CSResponseInviteRoomChat, EventResponseInviteRoomChat, middlewares...)
}

func EventResponseInviteRoomChat(client *socket_server.SocketClient, data *shared.SocketMessage) error {
	return nil
}
