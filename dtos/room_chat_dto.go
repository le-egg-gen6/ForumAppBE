package dtos

type RoomInfo struct {
	RoomID           int             `json:"roomID"`
	Name             string          `json:"name"`
	ParticipantInfos []SimpleUserDTO `json:"participantInfos"`
}
