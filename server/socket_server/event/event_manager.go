package event

import "forum/server/socket_server"

func RegisterEvent(router *socket_server.EventRouter) {
	RegisterEventLogin(router)
	RegisterEventAddFriend(router, AuthenticationEventMiddleware)
	RegisterEventGetFriendRequest(router, AuthenticationEventMiddleware)
	RegisterEventGetNotification(router, AuthenticationEventMiddleware)
	RegisterEventNewMessage(router, AuthenticationEventMiddleware)
	RegisterEventReactionMessage(router, AuthenticationEventMiddleware)
	RegisterEventCreateRoom(router, AuthenticationEventMiddleware)
	RegisterEventAddParticipantRoomChat(router, AuthenticationEventMiddleware)
	RegisterEventLeaveRoom(router, AuthenticationEventMiddleware)
	RegisterEventGetChatRoom(router, AuthenticationEventMiddleware)
	RegisterEventGetRoomMessage(router, AuthenticationEventMiddleware)
}
