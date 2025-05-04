package event

import (
	"forum/constant"
	"forum/server/socket_server"
	"forum/shared"
	"forum/utils"
)

func RegisterEventLogin(router *socket_server.EventRouter) {
	router.RegisterEventHandler(constant.CSLogin, EventLogin)
}

type CSLogin struct {
	Token string `json:"token"`
}

type SCLogin struct {
	Status int `json:"status"`
}

func EventLogin(client *socket_server.SocketClient, data *shared.SocketMessage) error {
	csLogin := utils.ConvertMessage[CSLogin](data)
	if csLogin == nil {
		utils.Disconnect(client)
	} else {
		utils.Send(client, constant.SCLogin, SCLogin{Status: 1})
	}
	return nil
}
