package sc

import "forum/dtos"

const StatusSuccess = 0
const StatusError = 1

type SCLogin struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type SCCreateRoom struct {
	Status   int           `json:"status"`
	Message  string        `json:"message"`
	RoomInfo dtos.RoomInfo `json:"roomInfo"`
}

type SCAddParticipantRoomChat struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type SCLeaveRoom struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type SCUpdateRoomInfo struct {
	Status   int           `json:"status"`
	Message  string        `json:"message"`
	RoomInfo dtos.RoomInfo `json:"roomInfo"`
}

type SCGetChatRoom struct {
	Status  int             `json:"status"`
	Message string          `json:"message"`
	Rooms   []dtos.RoomInfo `json:"rooms"`
}

type SCGetRoomMessage struct {
	Status   int                `json:"status"`
	Message  string             `json:"message"`
	Messages []dtos.MessageInfo `json:"messages"`
}

type SCNewMessage struct {
	Status      int              `json:"status"`
	Message     string           `json:"message"`
	RoomID      uint             `json:"roomID"`
	MessageInfo dtos.MessageInfo `json:"messageInfo"`
}

type SCReactionMessage struct {
	Status      int              `json:"status"`
	Message     string           `json:"message"`
	RoomID      uint             `json:"roomID"`
	MessageInfo dtos.MessageInfo `json:"messageInfo"`
}
