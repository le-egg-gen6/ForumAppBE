package event

import "forum/server/socket_server"

func RegisterEvent(router *socket_server.EventRouter) {
	RegisterEventLogin(router)
	RegisterEventAddFriend(router, AuthenticationEventMiddleware)
	RegisterEventGetFriendRequest(router, AuthenticationEventMiddleware)
	RegisterEventGetNotification(router, AuthenticationEventMiddleware)
	RegisterEventSendMessage(router, AuthenticationEventMiddleware)
	RegisterEventJoinRoom(router, AuthenticationEventMiddleware)
	RegisterEventLeaveRoom(router, AuthenticationEventMiddleware)
}
