package event

import (
	"forum/constant"
	"forum/server/socket_server"
	"forum/shared"
)

func RegisterEventNewMessage(router *socket_server.EventRouter, middleware ...socket_server.EventMiddlewareFunc) {
	router.RegisterEventHandler(constant.CSNewMessage, EventNewMessage, middleware...)
}

func EventNewMessage(client *socket_server.SocketClient, data *shared.SocketMessage) error {
	return nil
}
