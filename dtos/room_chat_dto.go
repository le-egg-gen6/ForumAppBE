package dtos

type RoomInfo struct {
	ID               uint            `json:"id"`
	Name             string          `json:"name"`
	Type             string          `json:"type"`
	ParticipantInfos []SimpleUserDTO `json:"participantInfos"`
}
