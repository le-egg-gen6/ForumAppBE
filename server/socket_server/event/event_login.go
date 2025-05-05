package event

import (
	"forum/3rd_party_service/redis_service"
	"forum/constant"
	"forum/server/socket_server"
	"forum/shared"
	"forum/utils"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
)

func RegisterEventLogin(router *socket_server.EventRouter, middleware ...socket_server.EventMiddlewareFunc) {
	router.RegisterEventHandler(constant.CSLogin, EventLogin, middleware...)
}

type CSLogin struct {
	Token string `json:"token"`
}

type SCLogin struct {
	Status int `json:"status"`
}

const Success = 0
const Failure = 1

func EventLogin(client *socket_server.SocketClient, data *shared.SocketMessage) error {
	csLogin := utils.ConvertMessage[CSLogin](data)
	if csLogin == nil {
		utils.Disconnect(client)
	} else {
		tokenStr := csLogin.Token
		tokenUsed, _ := redis_service.Get[bool](tokenStr)
		if tokenUsed {
			SendFailure(client)
			return nil
		}
		jwtToken, err := utils.ValidateToken(tokenStr)
		if err != nil {
			SendFailure(client)
			return nil
		}
		claims, ok := jwtToken.Claims.(*jwt.RegisteredClaims)
		if !ok {
			SendFailure(client)
			return nil
		}
		userIDToken, err := strconv.ParseUint(claims.Subject, 10, 64)
		if err != nil {
			SendFailure(client)
			return nil
		}
		if userIDToken != client.UserID {
			SendFailure(client)
			return nil
		}
		SendSuccess(client)
	}
	return nil
}

func SendFailure(client *socket_server.SocketClient) {
	utils.Send(client, constant.SCLogin, SCLogin{Status: Failure})
	utils.Disconnect(client)
}

func SendSuccess(client *socket_server.SocketClient) {
	utils.Send(client, constant.SCLogin, SCLogin{Status: Success})
	client.Hub.AuthorizeSocketConnection(client)
}
