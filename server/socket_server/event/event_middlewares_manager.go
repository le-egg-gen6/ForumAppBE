package event

import (
	"forum/server/socket_server"
	"forum/shared"
)

func AuthenticationEventMiddleware(client *socket_server.SocketClient, data *shared.SocketMessage) bool {
	return client.Authorized
}
