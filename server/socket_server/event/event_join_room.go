package event

import (
	"forum/constant"
	"forum/server/socket_server"
	"forum/shared"
)

func RegisterEventJoinRoom(router *socket_server.EventRouter) {
	router.RegisterEventHandler(constant.CSJoinRoom, EventJoinRoom)
}

func EventJoinRoom(client *socket_server.SocketClient, data *shared.SocketMessage) error {
	return nil
}
