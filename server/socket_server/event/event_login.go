package event

import (
	"forum/3rd_party_service/redis_service"
	"forum/constant"
	"forum/server/socket_server"
	"forum/server/socket_server/message/cs"
	"forum/server/socket_server/message/sc"
	"forum/shared"
	"forum/utils"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
)

func RegisterEventLogin(router *socket_server.EventRouter, middleware ...socket_server.EventMiddlewareFunc) {
	router.RegisterEventHandler(constant.CSLogin, EventLogin, middleware...)
}

func EventLogin(client *socket_server.SocketClient, data *shared.SocketMessage) error {
	csLogin := utils.ConvertMessage[cs.CSLogin](data)
	if csLogin == nil {
		SendLoginFailure(client, "Bad request")
		return nil
	}
	tokenStr := csLogin.Token
	tokenUsed, _ := redis_service.Get[bool](tokenStr)
	if tokenUsed {
		SendLoginFailure(client, "Token invalid")
		return nil
	}
	jwtToken, err := utils.ValidateToken(tokenStr)
	if err != nil {
		SendLoginFailure(client, "Token invalid")
		return nil
	}
	claims, ok := jwtToken.Claims.(*jwt.RegisteredClaims)
	if !ok {
		SendLoginFailure(client, "Token invalid")
		return nil
	}
	userIDToken, err := strconv.Atoi(claims.Subject)
	if err != nil {
		SendLoginFailure(client, "Token invalid")
		return nil
	}
	if uint(userIDToken) != client.UserID {
		SendLoginFailure(client, "Token invalid")
		return nil
	}
	SendLoginSuccess(client)
	return nil
}

func SendLoginFailure(client *socket_server.SocketClient, message string) {
	utils.Send(client, constant.SCLogin, sc.SCLogin{Status: sc.StatusError, Message: message})
	utils.Disconnect(client)
}

func SendLoginSuccess(client *socket_server.SocketClient) {
	utils.Send(client, constant.SCLogin, sc.SCLogin{Status: sc.StatusSuccess})
	client.Hub.AuthorizeSocketConnection(client)
}
