package event

import (
	"forum/constant"
	"forum/server/socket_server"
	"forum/shared"
)

func RegisterEventLeaveRoom(router *socket_server.EventRouter, middleware ...socket_server.EventMiddlewareFunc) {
	router.RegisterEventHandler(constant.CSLeaveRoom, EventLeaveRoom, middleware...)
}

func EventLeaveRoom(client *socket_server.SocketClient, data *shared.SocketMessage) error {
	return nil
}
