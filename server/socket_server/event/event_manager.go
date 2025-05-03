package event

import "forum/server/socket_server"

func RegisterEvent(router *socket_server.EventRouter) {
	RegisterEventLogin(router)
	RegisterEventAddFriend(router)
	RegisterEventGetFriendRequest(router)
	RegisterEventGetNotification(router)
	RegisterEventSendMessage(router)
}
