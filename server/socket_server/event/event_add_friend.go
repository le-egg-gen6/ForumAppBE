package event

import (
	"forum/constant"
	"forum/server/socket_server"
	"forum/shared"
)

func RegisterEventAddFriend(router *socket_server.EventRouter) {
	router.RegisterEventHandler(constant.CSAddFriend, EventAddFriend)
}

func EventAddFriend(client *socket_server.SocketClient, data *shared.SocketMessage) error {
	return nil
}
