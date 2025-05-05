package event

import (
	"forum/constant"
	"forum/server/socket_server"
	"forum/shared"
)

func RegisterEventReactionMessage(router *socket_server.EventRouter, middleware ...socket_server.EventMiddlewareFunc) {
	router.RegisterEventHandler(constant.CSReactionMessage, ReactionMessage, middleware...)
}

func ReactionMessage(client *socket_server.SocketClient, data *shared.SocketMessage) error {
	return nil
}
