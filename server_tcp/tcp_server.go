package server_tcp

import socketio "github.com/googollee/go-socket.io"

func NewSocketServer() *socketio.Server {
	server := socketio.NewServer(nil)
	RegisterAllHandler(server)
	return server
}
