package socket_server

import (
	"forum/logger"
	"sync"
)

type EventHandlerFunc func(client *SocketClient, data interface{}) error

type EventRouter struct {
	EventHandlers map[string]EventHandlerFunc
	rwMutex       sync.RWMutex
	Hub           *Hub
}

func InitializeEventRouter(hub *Hub) *EventRouter {
	return &EventRouter{
		EventHandlers: make(map[string]EventHandlerFunc),
		Hub:           hub,
	}
}

func (er *EventRouter) RegisterEventHandler(name string, handler EventHandlerFunc) {
	er.rwMutex.Lock()
	defer er.rwMutex.Unlock()
	er.EventHandlers[name] = handler
	logger.GetLogInstance().Info("Successfully registered event: " + name)
}

func (er *EventRouter) GetEventHandler(name string) (EventHandlerFunc, bool) {
	er.rwMutex.RLock()
	defer er.rwMutex.RUnlock()
	handler, found := er.EventHandlers[name]
	return handler, found
}
