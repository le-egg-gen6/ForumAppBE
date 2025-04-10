package tcp

import (
	socketio "github.com/googollee/go-socket.io"
)

type ChatHandler struct{}

func NewChatHandler() *ChatHandler {
	return &ChatHandler{}
}

func (h *ChatHandler) Name() string {
	return "chat_handler"
}

func (h *ChatHandler) Register(server *socketio.Server) {

}
