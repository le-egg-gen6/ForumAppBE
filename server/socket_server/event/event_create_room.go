package event

import (
	"forum/constant"
	"forum/server/socket_server"
	"forum/shared"
)

func RegisterEventCreateRoom(router *socket_server.EventRouter, middleware ...socket_server.EventMiddlewareFunc) {
	router.RegisterEventHandler(constant.CSJoinRoom, EventCreateRoom, middleware...)
}

func EventCreateRoom(client *socket_server.SocketClient, data *shared.SocketMessage) error {
	return nil
}
