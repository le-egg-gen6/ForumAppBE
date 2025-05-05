package event

import (
	"forum/constant"
	"forum/server/socket_server"
	"forum/shared"
)

func RegisterEventGetFriendRequest(router *socket_server.EventRouter, middleware ...socket_server.EventMiddlewareFunc) {
	router.RegisterEventHandler(constant.CSGetFriendRequest, EventGetFriendRequest, middleware...)
}

func EventGetFriendRequest(client *socket_server.SocketClient, data *shared.SocketMessage) error {
	return nil
}
