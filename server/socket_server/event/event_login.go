package event

import (
	"fmt"
	"forum/constant"
	"forum/server/socket_server"
	"forum/shared"
	"forum/utils"
)

func RegisterEventLogin(router *socket_server.EventRouter) {
	router.RegisterEventHandler(constant.CSLogin, EventLogin)
}

func EventLogin(client *socket_server.SocketClient, data interface{}) error {
	utils.Send(client, shared.SocketMessage{
		Name: constant.SCLogin,
		Data: fmt.Sprintf("Server resp: %v", data),
	})
	return nil
}
