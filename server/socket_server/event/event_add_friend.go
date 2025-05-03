package event

import (
	"forum/constant"
	"forum/server/socket_server"
)

func RegisterEventAddFriend(router *socket_server.EventRouter) {
	router.RegisterEventHandler(constant.CSAddFriend, EventAddFriend)
}

func EventAddFriend(client *socket_server.SocketClient, data interface{}) error {
	return nil
}
