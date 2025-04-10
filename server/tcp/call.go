package tcp

import socketio "github.com/googollee/go-socket.io"

type CalllHandler struct{}

func NewCallHandler() *ChatHandler {
	return &ChatHandler{}
}

func (h *CalllHandler) Name() string {
	return "call_handler"
}

func (h *CalllHandler) Register(server *socketio.Server) {

}
