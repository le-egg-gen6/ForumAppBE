package event

import (
	"forum/constant"
	"forum/server/socket_server"
)

func RegisterEventSendMessage(router *socket_server.EventRouter) {
	router.RegisterEventHandler(constant.CSSendMessage, EventSendMessage)
}

func EventSendMessage(client *socket_server.SocketClient, data interface{}) error {
	return nil
}
