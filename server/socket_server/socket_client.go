package socket_server

import (
	"fmt"
	"github.com/gorilla/websocket"
	"sync"
)

type SocketClient struct {
	ID         string
	UserID     uint64
	Authorized bool
	Conn       *websocket.Conn
	mutex      sync.Mutex
	Hub        *Hub
}

func (c *SocketClient) Send(data interface{}) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.Conn == nil {
		return fmt.Errorf("connection is nil for client %s", c.ID)
	}
	return c.Conn.WriteJSON(data)
}
