package tcp

import socketio "github.com/googollee/go-socket.io"

import (
	"fmt"
)

type ConnectionHandler struct{}

func NewConnectionHandler() *ConnectionHandler {
	return &ConnectionHandler{}
}

func (h *ConnectionHandler) Name() string {
	return "connection_handler"
}

func (h *ConnectionHandler) Register(server *socketio.Server) {

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("Client connected:", s.ID())
		return nil
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("Client disconnected:", s.ID(), reason)
	})

}
