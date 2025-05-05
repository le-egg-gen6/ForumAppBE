package sc

import "forum/dtos"

type SCLogin struct {
	Status int `json:"status"`
}

const LoginSuccess = 0
const LoginFailure = 1

type SCCreateRoom struct {
	Status           int                  `json:"status"`
	RoomID           uint64               `json:"roomID"`
	Name             string               `json:"name"`
	ParticipantInfos []dtos.SimpleUserDTO `json:"participantInfos"`
}

type SCLeaveRoom struct {
	Status int `json:"status"`
}

type SCGetChatRoom struct {
	Rooms []dtos.RoomInfo `json:"rooms"`
}

type SCGetRoomMessage struct {
	Messages []dtos.MessageInfo `json:"messages"`
}

type SCNewMessage struct {
	RoomID      uint64           `json:"roomID"`
	MessageInfo dtos.MessageInfo `json:"messageInfo"`
}

type SCReactionMessage struct {
	RoomID      uint64           `json:"roomID"`
	MessageInfo dtos.MessageInfo `json:"messageInfo"`
}

type SCGetNotification struct {
}

type SCGetFriendRequest struct {
}

type SCNewNotification struct {
}
