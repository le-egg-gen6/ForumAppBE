package event

import (
	"forum/constant"
	"forum/server/socket_server"
	"forum/shared"
)

func RegisterEventGetNotification(router *socket_server.EventRouter, middleware ...socket_server.EventMiddlewareFunc) {
	router.RegisterEventHandler(constant.CSGetNotification, EventGetNotification, middleware...)
}

func EventGetNotification(client *socket_server.SocketClient, data *shared.SocketMessage) error {
	return nil
}
