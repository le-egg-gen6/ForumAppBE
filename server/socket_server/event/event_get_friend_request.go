package event

import (
	"forum/constant"
	"forum/server/socket_server"
)

func RegisterEventGetFriendRequest(router *socket_server.EventRouter) {
	router.RegisterEventHandler(constant.CSGetFriendRequest, EventGetFriendRequest)
}

func EventGetFriendRequest(client *socket_server.SocketClient, data interface{}) error {
	return nil
}
