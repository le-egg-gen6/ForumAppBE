package event

import (
	"forum/constant"
	"forum/server/socket_server"
)

func RegisterEventLeaveRoom(router *socket_server.EventRouter) {
	router.RegisterEventHandler(constant.CSLeaveRoom, EventLeaveRoom)
}

func EventLeaveRoom(client *socket_server.SocketClient, data interface{}) error {
	return nil
}
