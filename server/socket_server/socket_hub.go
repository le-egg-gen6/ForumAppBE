package socket_server

import (
	"fmt"
	"forum/logger"
	"github.com/gorilla/websocket"
	"sync"
)

type Hub struct {
	Clients        map[*websocket.Conn]*SocketClient
	ClientByUserID map[uint]*SocketClient
	Register       chan *SocketClient
	Unregister     chan *SocketClient
	rwMutex        sync.RWMutex
}

func InitializeNewHub() *Hub {
	return &Hub{
		Clients:        make(map[*websocket.Conn]*SocketClient),
		ClientByUserID: make(map[uint]*SocketClient),
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

func (h *Hub) GetClientByUserID(userID uint) *SocketClient {
	h.rwMutex.RLock()
	defer h.rwMutex.RUnlock()
	socketClient, found := h.ClientByUserID[userID]
	if !found {
		return nil
	}
	return socketClient
}

func (h *Hub) GetClientsByUserIDs(userIDs []uint) []*SocketClient {
	h.rwMutex.RLock()
	defer h.rwMutex.RUnlock()
	var clients []*SocketClient
	for _, userID := range userIDs {
		if client, found := h.ClientByUserID[userID]; found {
			clients = append(clients, client)
		}
	}
	return clients
}

func (h *Hub) AuthorizeSocketConnection(client *SocketClient) {
	h.rwMutex.Lock()
	defer h.rwMutex.Unlock()
	client.Authorized = true
}

func (h *Hub) Close() {
	h.rwMutex.Lock()
	defer h.rwMutex.Unlock()
	for conn, client := range h.Clients {
		if client != nil && client.Conn != nil {
			err := client.Conn.Close()
			if err != nil {
				logger.GetLogInstance().Error(fmt.Sprintf(
					"Failed to close connection for client ID: %s, UserID: %d, Addr: %s, Error: %v",
					client.ID,
					client.UserID,
					client.Conn.RemoteAddr(),
					err,
				))
			}
			delete(h.ClientByUserID, client.UserID)
		}
		delete(h.Clients, conn)
	}
	if len(h.Clients) > 0 {
		logger.GetLogInstance().Warn(fmt.Sprintf("Still have connection remaining in map clients by connection: %d", len(h.Clients)))
	}
	if len(h.ClientByUserID) > 0 {
		logger.GetLogInstance().Warn(fmt.Sprintf("Still have connection remaining in map clients by user id: %d", len(h.ClientByUserID)))
	}
	close(h.Register)
	close(h.Unregister)
}
