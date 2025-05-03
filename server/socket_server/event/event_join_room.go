package event

import (
	"forum/constant"
	"forum/server/socket_server"
)

func RegisterEventJoinRoom(router *socket_server.EventRouter) {
	router.RegisterEventHandler(constant.CSJoinRoom, EventJoinRoom)
}

func EventJoinRoom(client *socket_server.SocketClient, data interface{}) error {
	return nil
}
