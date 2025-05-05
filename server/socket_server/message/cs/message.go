package cs

type CSLogin struct {
	Token string `json:"token"`
}

type CSCreateRoom struct {
	Name           string   `json:"name"`
	ParticipantIDs []uint64 `json:"participantIDs"`
}

type CSLeaveRoom struct {
	RoomID uint64 `json:"roomID"`
}

type CSGetChatRoom struct {
}

type CSGetRoomMessage struct {
	RoomID uint64 `json:"roomID"`
}

type CSNewMessage struct {
	RoomID uint64 `json:"roomID"`
	Body   string `json:"body"`
}

type CSReactionMessage struct {
	RoomID    uint64 `json:"roomID"`
	MessageID uint64 `json:"messageID"`
	Type      string `json:"type"`
}

type CSAddFriend struct {
	FriendID uint64 `json:"friendID"`
}

type CSGetNotification struct {
}

type CSGetFriendRequest struct {
}
