package dtos

type RoomInfo struct {
	RoomID           uint64          `json:"roomID"`
	Name             string          `json:"name"`
	Type             string          `json:"type"`
	ParticipantInfos []SimpleUserDTO `json:"participantInfos"`
}
