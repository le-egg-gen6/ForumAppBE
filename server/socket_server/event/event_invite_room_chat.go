package event

import (
	"forum/constant"
	"forum/server/socket_server"
	"forum/shared"
)

func RegisterEventInviteRoomChat(router *socket_server.EventRouter, middlewares ...socket_server.EventMiddlewareFunc) {
	router.RegisterEventHandler(constant.CSInviteRoomChat, EventInviteRoomChat, middlewares...)
}

func EventInviteRoomChat(client *socket_server.SocketClient, data *shared.SocketMessage) error {
	return nil
}
