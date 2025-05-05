package socket_server

import (
	"encoding/json"
	"fmt"
	"forum/logger"
	"forum/shared"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"net/http"
	"strconv"
)

type SocketServer struct {
	Config *SocketServerConfig
	Hub    *Hub
	Router *EventRouter
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var Server *SocketServer

func InitializeSocketServer() {
	cfg, err := LoadSocketServerConfig()
	if err != nil {
		panic("TCP Server configuration not found")
	}
	hub := InitializeNewHub()
	eventRouter := InitializeEventRouter(hub)
	SocketIOServer := &SocketServer{
		Config: cfg,
		Hub:    hub,
		Router: eventRouter,
	}
	Server = SocketIOServer
}

func GetSocketServer() *SocketServer {
	return Server
}

func HandleConnection(hub *Hub, router *EventRouter, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.GetLogInstance().Warn(fmt.Sprintf("Upgrade failed: %v", err))
		return
	}
	userIDStr := r.URL.Query().Get("id")
	var userID int64
	if userIDStr != "" {
		userID, err = strconv.ParseInt(userIDStr, 10, 64)
		if err != nil || userID < 0 {
			logger.GetLogInstance().Error(fmt.Sprintf("Parse user id failed: %v", err))
			return
		}
	}
	client := &SocketClient{
		ID:         uuid.NewString(),
		UserID:     uint64(userID),
		Authorized: false,
		Conn:       conn,
		Hub:        hub,
	}

	hub.Register <- client

	go ReadPump(client, router)
}

func ReadPump(client *SocketClient, router *EventRouter) {
	conn := client.Conn
	defer func() {
		client.Hub.Unregister <- client
	}()

	for {
		messageType, messageBytes, err := conn.ReadMessage()
		if err == nil {
			if messageType != websocket.TextMessage {
				logger.GetLogInstance().Info(fmt.Sprintf(
					"Client ID: %s, UserID: %d, Addr: %s send non-text message",
					client.ID,
					client.UserID,
					client.Conn.RemoteAddr(),
				))
				continue
			}
			var msg shared.SocketMessage
			err = json.Unmarshal(messageBytes, &msg)
			if err != nil {
				logger.GetLogInstance().Error(fmt.Sprintf(
					"Client ID: %s, UserID: %d, Addr: %s send invalid message format. Raw: %s",
					client.ID,
					client.UserID,
					client.Conn.RemoteAddr(),
					string(messageBytes),
				))
				continue
			}
			middlewares, found := router.GetMiddlewares(msg.Name)
			passMiddleware := true
			if found {
				for _, middleware := range middlewares {
					pass := middleware(client, &msg)
					passMiddleware = passMiddleware && pass
				}
			}
			if !passMiddleware {
				break
			}
			handler, found := router.GetEventHandler(msg.Name)
			if !found {
				logger.GetLogInstance().Error(fmt.Sprintf(
					"No handler found for event: %s. Client ID: %s, UserID: %d, Addr: %s.",
					msg.Name,
					client.ID,
					client.UserID,
					client.Conn.RemoteAddr(),
				))
				continue
			}
			err = handler(client, &msg)
			if err != nil {
				logger.GetLogInstance().Error(fmt.Sprintf(
					"Error executing handler for event '%s'. Client ID: %s, UserID: %d, Addr: %s send message failed. Raw: %s",
					msg.Name,
					client.ID,
					client.UserID,
					client.Conn.RemoteAddr(),
					string(messageBytes),
				))
			}
		} else {
			if websocket.IsUnexpectedCloseError(err,
				websocket.CloseGoingAway, websocket.CloseAbnormalClosure, websocket.CloseNoStatusReceived) {
				logger.GetLogInstance().Info(fmt.Sprintf(
					"Client ID: %s, UserID: %d, Addr: %s disconnected unexpectedly: %v",
					client.ID,
					client.UserID,
					conn.RemoteAddr(),
					err,
				))
			} else {
				logger.GetLogInstance().Error(fmt.Sprintf(
					"Client ID: %s, UserID: %d, Addr: %s reading error: %v ",
					client.ID,
					client.UserID,
					conn.RemoteAddr(),
					err,
				))
			}
			break
		}
	}
}

func (ss *SocketServer) RegisterEventHandler(name string, handler EventHandlerFunc) {
	ss.Router.RegisterEventHandler(name, handler)
}

func (ss *SocketServer) Run() error {
	go ss.Hub.Run()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		HandleConnection(ss.Hub, ss.Router, w, r)
	})
	addr := fmt.Sprintf(":%d", ss.Config.Port)
	return http.ListenAndServe(addr, nil)
}

func (ss *SocketServer) Close() error {
	if ss == nil || ss.Hub == nil {
		return fmt.Errorf("socket server or hub is not initialized")
	}
	ss.Hub.Close()
	return nil
}
