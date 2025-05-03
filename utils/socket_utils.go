package utils

import "forum/server/socket_server"

func Send(client *socket_server.SocketClient, data interface{}) {
	_ = client.Send(data)
}

func Broadcast(clients []*socket_server.SocketClient, data interface{}) {
	for _, client := range clients {
		_ = client.Send(data)
	}
}
