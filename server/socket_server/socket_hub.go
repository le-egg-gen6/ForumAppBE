package socket_server

import (
	"fmt"
	"forum/logger"
	"github.com/gorilla/websocket"
	"sync"
)

type Hub struct {
	Clients        map[*websocket.Conn]*SocketClient
	ClientByUserID map[uint64]*SocketClient
	Register       chan *SocketClient
	Unregister     chan *SocketClient
	rwMutex        sync.RWMutex
}

func InitializeNewHub() *Hub {
	return &Hub{
		Clients:        make(map[*websocket.Conn]*SocketClient),
		ClientByUserID: make(map[uint64]*SocketClient),
		Register:       make(chan *SocketClient),
		Unregister:     make(chan *SocketClient),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			if client == nil || client.Conn == nil {
				logger.GetLogInstance().Warn("Attempt to Register nil client/connection")
				continue
			}
			h.rwMutex.Lock()
			h.Clients[client.Conn] = client
			h.ClientByUserID[client.UserID] = client
			h.rwMutex.Unlock()
			logger.GetLogInstance().Info(fmt.Sprintf(
				"Register new client: ID: %s, UserID: %d, Addr: %s",
				client.ID,
				client.UserID,
				client.Conn.RemoteAddr(),
			))
		case client := <-h.Unregister:
			if client == nil || client.Conn == nil {
				logger.GetLogInstance().Warn("Attempt to Unregister nil client/connection")
				continue
			}
			conn := client.Conn
			h.rwMutex.Lock()
			if registeredClient, ok := h.Clients[conn]; ok {
				delete(h.Clients, conn)
				delete(h.ClientByUserID, registeredClient.UserID)
				logger.GetLogInstance().Info(fmt.Sprintf(
					"Unregister client: ID: %s, UserID: %d, Addr: %s",
					client.ID,
					registeredClient.UserID,
					conn.RemoteAddr(),
				))
			}
			err := conn.Close()
			if err != nil {
				logger.GetLogInstance().Error("Error while close websocket connection" + err.Error())
			}
			h.rwMutex.Unlock()

		}
	}
}

func (h *Hub) GetClientByUserID(userID uint64) *SocketClient {
	h.rwMutex.RLock()
	defer h.rwMutex.RUnlock()
	socketClient, found := h.ClientByUserID[userID]
	if !found {
		return nil
	}
	return socketClient
}
