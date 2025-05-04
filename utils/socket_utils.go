package utils

import (
	"encoding/json"
	"forum/server/socket_server"
	"forum/shared"
)

func Send(client *socket_server.SocketClient, name string, data interface{}) {
	socketMsg := shared.SocketMessage{
		Name: name,
		Data: data,
	}
	_ = client.Send(socketMsg)
}

func Broadcast(clients []*socket_server.SocketClient, name string, data interface{}) {
	socketMsg := shared.SocketMessage{
		Name: name,
		Data: data,
	}
	for _, client := range clients {
		_ = client.Send(socketMsg)
	}
}

func Disconnect(client *socket_server.SocketClient) {
	if client == nil || client.Conn == nil {
		return
	}
	client.Hub.Unregister <- client
}

func ConvertMessage[T any](socketMsg *shared.SocketMessage) *T {
	raw, err := json.Marshal(socketMsg.Data)
	if err != nil {
		return nil
	}
	var csMsg T
	err = json.Unmarshal(raw, &csMsg)
	if err != nil {
		return nil
	}
	return &csMsg
}

func IsUserOnline(userID uint64) bool {
	_, found := socket_server.GetSocketServer().Hub.ClientByUserID[userID]
	return found
}
