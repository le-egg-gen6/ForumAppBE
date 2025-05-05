package event

import (
	"forum/constant"
	"forum/server/socket_server"
	"forum/shared"
)

func RegisterEventGetChatRoom(router *socket_server.EventRouter, middleware ...socket_server.EventMiddlewareFunc) {
	router.RegisterEventHandler(constant.CSGetChatRoom, EventGetChatRoom, middleware...)
}

func EventGetChatRoom(client *socket_server.SocketClient, data *shared.SocketMessage) error {
	return nil
}
