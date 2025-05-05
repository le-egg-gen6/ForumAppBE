package event

import (
	"forum/constant"
	"forum/server/socket_server"
	"forum/shared"
)

type CSJoinRoom struct {
	RoomID uint64 `json:"roomID"`
}

func RegisterEventJoinRoom(router *socket_server.EventRouter, middleware ...socket_server.EventMiddlewareFunc) {
	router.RegisterEventHandler(constant.CSJoinRoom, EventJoinRoom, middleware...)
}

func EventJoinRoom(client *socket_server.SocketClient, data *shared.SocketMessage) error {
	return nil
}
