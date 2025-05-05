package event

import (
	"forum/constant"
	"forum/server/socket_server"
	"forum/shared"
)

func RegisterEventSendMessage(router *socket_server.EventRouter, middleware ...socket_server.EventMiddlewareFunc) {
	router.RegisterEventHandler(constant.CSSendMessage, EventSendMessage, middleware...)
}

func EventSendMessage(client *socket_server.SocketClient, data *shared.SocketMessage) error {
	return nil
}
