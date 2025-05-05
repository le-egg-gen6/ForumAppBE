package cs

type CSLogin struct {
	Token string `json:"token"`
}

type CSJoinRoom struct {
	RoomID uint64 `json:"roomID"`
}
