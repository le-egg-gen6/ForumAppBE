package tcp

import socketio "github.com/googollee/go-socket.io"

type SocketHandler interface {
	Name() string
	Register(server *socketio.Server)
}

var handlers = make(map[string]SocketHandler)

func RegisterHandler(handler SocketHandler) {
	handlers[handler.Name()] = handler
}

func RegisterAllHandler(server *socketio.Server) {
	RegisterHandler(NewConnectionHandler())
	RegisterHandler(NewChatHandler())
	RegisterHandler(NewCallHandler())

	for _, handler := range handlers {
		handler.Register(server)
	}

}
