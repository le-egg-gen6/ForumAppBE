package event

import (
	"forum/constant"
	"forum/models"
	"forum/server/socket_server"
	"forum/server/socket_server/message/sc"
	"forum/utils"
)

func BroadcastRoomNewMessage(
	client *socket_server.SocketClient,
	room *models.RoomChat,
	message *models.RoomMessage,
	author *models.User,
) {
	userIDs := make([]uint, 0)
	for _, user := range room.Users {
		userIDs = append(userIDs, user.ID)
	}
	roomMembersSocketConn := client.Hub.GetClientsByUserIDs(userIDs)
	scMessage := sc.SCNewMessage{
		RoomID:      room.ID,
		Status:      sc.StatusSuccess,
		MessageInfo: *utils.ConvertToRoomMessageInfo(message, author),
	}
	utils.Broadcast(roomMembersSocketConn, constant.SCNewMessage, scMessage)
}

func BroadcastUpdateRoomInfo(
	client *socket_server.SocketClient,
	room *models.RoomChat,
) {
	userIDs := make([]uint, 0)
	for _, user := range room.Users {
		userIDs = append(userIDs, user.ID)
	}
	roomMembersSocketConn := client.Hub.GetClientsByUserIDs(userIDs)
	scMessage := sc.SCUpdateRoomInfo{
		Status:   sc.StatusSuccess,
		RoomInfo: *utils.ConvertToRoomInfo(room),
	}
	utils.Broadcast(roomMembersSocketConn, constant.SCUpdateRoomInfo, scMessage)
}
