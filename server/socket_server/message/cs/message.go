package cs

type CSLogin struct {
	Token string `json:"token"`
}

type CSCreateRoom struct {
	Name           string `json:"name"`
	ParticipantIDs []uint `json:"participantIDs"`
}

type CSAddParticipantRoomChat struct {
	RoomID        uint `json:"roomID"`
	ParticipantID uint `json:"participantID"`
}

type CSLeaveRoom struct {
	RoomID uint `json:"roomID"`
}

type CSGetChatRoom struct {
}

type CSGetRoomMessage struct {
	RoomID uint `json:"roomID"`
}

type CSNewMessage struct {
	RoomID uint   `json:"roomID"`
	Body   string `json:"body"`
}

type CSReactionMessage struct {
	RoomID    uint   `json:"roomID"`
	MessageID uint   `json:"messageID"`
	Type      string `json:"type"`
}

type CSAddFriend struct {
	FriendID uint `json:"friendID"`
}

type CSGetNotification struct {
}

type CSGetFriendRequest struct {
}
