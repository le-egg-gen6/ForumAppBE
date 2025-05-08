package socket_server

import (
	"forum/logger"
	"forum/shared"
	"sync"
)

type EventMiddlewareFunc func(client *SocketClient, data *shared.SocketMessage) bool

type EventHandlerFunc func(client *SocketClient, data *shared.SocketMessage) error

type EventRouter struct {
	EventHandlers       map[string]EventHandlerFunc
	EventMiddlewareFunc map[string][]EventMiddlewareFunc
	rwMutex             sync.RWMutex
	Hub                 *Hub
}

func InitializeEventRouter(hub *Hub) *EventRouter {
	return &EventRouter{
		EventHandlers:       make(map[string]EventHandlerFunc),
		EventMiddlewareFunc: make(map[string][]EventMiddlewareFunc),
		Hub:                 hub,
	}
}

func (er *EventRouter) RegisterEventHandler(name string, handler EventHandlerFunc, middlewares ...EventMiddlewareFunc) {
	er.rwMutex.Lock()
	defer er.rwMutex.Unlock()
	er.EventHandlers[name] = handler
	er.EventMiddlewareFunc[name] = middlewares
	logger.GetLogInstance().Info("Successfully registered event: " + name)
}

func (er *EventRouter) GetMiddlewares(name string) ([]EventMiddlewareFunc, bool) {
	er.rwMutex.RLock()
	defer er.rwMutex.RUnlock()
	middlewares, found := er.EventMiddlewareFunc[name]
	return middlewares, found
}

func (er *EventRouter) GetEventHandler(name string) (EventHandlerFunc, bool) {
	er.rwMutex.RLock()
	defer er.rwMutex.RUnlock()
	handler, found := er.EventHandlers[name]
	return handler, found
}
