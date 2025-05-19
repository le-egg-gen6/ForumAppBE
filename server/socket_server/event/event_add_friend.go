package event

import (
	"forum/constant"
	"forum/server/socket_server"
	"forum/server/socket_server/message/cs"
	"forum/server/socket_server/message/sc"
	"forum/shared"
	"forum/utils"
)

func RegisterEventAddFriend(router *socket_server.EventRouter, middleware ...socket_server.EventMiddlewareFunc) {
	router.RegisterEventHandler(constant.CSAddFriend, EventAddFriend, middleware...)
}

func EventAddFriend(client *socket_server.SocketClient, data *shared.SocketMessage) error {
	csAddFriend := utils.ConvertMessage[cs.CSAddFriend](data)
	if csAddFriend == nil {
		AddFriendFailure(client, "Bad request")
		return nil
	}
	
}

func AddFriendFailure(client *socket_server.SocketClient, message string) {
	utils.Send(client, constant.SCAddFriend, sc.SCAddFriend{Status: sc.StatusError, Message: message})
}

func AddFriendSuccess(client *socket_server.SocketClient) {
	utils.Send(client, constant.SCAddFriend, sc.SCAddFriend{Status: sc.StatusSuccess})
}
