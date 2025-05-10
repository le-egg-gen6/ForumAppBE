package sc

import "forum/dtos"

type SCLogin struct {
	Status int `json:"status"`
}

const LoginSuccess = 0
const LoginFailure = 1

type SCCreateRoom struct {
	Status   int           `json:"status"`
	RoomInfo dtos.RoomInfo `json:"roomInfo"`
}

const RoomCreatedSuccess = 0
const RoomCreateFailed = 1

type SCLeaveRoom struct {
	Status int `json:"status"`
}

const RoomLeaveSuccess = 0
const RoomLeaveFailed = 1

type SCGetChatRoom struct {
	Rooms []dtos.RoomInfo `json:"rooms"`
}

type SCGetRoomMessage struct {
	Messages []dtos.MessageInfo `json:"messages"`
}

type SCNewMessage struct {
	RoomID      uint             `json:"roomID"`
	Status      int              `json:"status"`
	MessageInfo dtos.MessageInfo `json:"messageInfo"`
}

const SendNewMessageSuccess = 0
const SendNewMessageFailed = 1

type SCReactionMessage struct {
	RoomID      uint             `json:"roomID"`
	MessageInfo dtos.MessageInfo `json:"messageInfo"`
}

type SCGetNotification struct {
}

type SCGetFriendRequest struct {
}

type SCNewNotification struct {
}
