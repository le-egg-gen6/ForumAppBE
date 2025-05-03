package event

import (
	"forum/constant"
	"forum/server/socket_server"
)

func RegisterEventGetNotification(router *socket_server.EventRouter) {
	router.RegisterEventHandler(constant.CSGetNotification, EventGetNotification)
}

func EventGetNotification(client *socket_server.SocketClient, data interface{}) error {
	return nil
}
