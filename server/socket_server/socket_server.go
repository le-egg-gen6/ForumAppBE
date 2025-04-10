package socket_server

import (
	"fmt"
	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"net/http"
	"sync"
	"time"
)

type SocketIOServer struct {
	Config     *SocketServerConfig
	Server     *socketio.Server
	httpServer *http.Server
	clients    map[string]socketio.Conn
	clientsMux sync.RWMutex
}

var Server *SocketIOServer

func InitializeSocketIOServer() {
	cfg, err := LoadSocketServerConfig()
	if err != nil {
		panic("TCP Server configuration not found")
	}
	server := socketio.NewServer(&engineio.Options{
		PingTimeout:  time.Duration(cfg.PingTimeoutSec) * time.Second,
		PingInterval: time.Duration(cfg.PingIntervalSec) * time.Second,
	})
	mux := http.NewServeMux()
	mux.Handle("/socket.io/", server)
	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: mux,
	}

	Server = &SocketIOServer{
		Config:     cfg,
		Server:     server,
		httpServer: httpServer,
	}
}

func GetSocketServer() *SocketIOServer {
	return Server
}

func (s *SocketIOServer) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *SocketIOServer) RegisterEvent(event string, handler func(socketio.Conn, interface{})) {
	s.Server.OnEvent("/", event, handler)
}

func (s *SocketIOServer) SendToClient(clientID, event string, data interface{}) bool {
	s.clientsMux.RLock()
	defer s.clientsMux.RUnlock()

	if client, exists := s.clients[clientID]; exists {
		client.Emit(event, data)
		return true
	}

	return false
}

func (s *SocketIOServer) GetAllClients() []string {
	s.clientsMux.RLock()
	defer s.clientsMux.RUnlock()

	clients := make([]string, 0, len(s.clients))
	for clientID := range s.clients {
		clients = append(clients, clientID)
	}

	return clients
}

func (s *SocketIOServer) Close() error {
	if err := s.Server.Close(); err != nil {
		return err
	}

	return s.httpServer.Close()
}
