package event

import (
	"forum/constant"
	"forum/server/socket_server"
	"forum/shared"
)

func RegisterEventGetNotification(router *socket_server.EventRouter) {
	router.RegisterEventHandler(constant.CSGetNotification, EventGetNotification)
}

func EventGetNotification(client *socket_server.SocketClient, data *shared.SocketMessage) error {
	return nil
}
