package dtos

import "time"

type MessageInfo struct {
	MessageID uint64        `json:"messageID"`
	Body      string        `json:"body"`
	Type      string        `json:"type"`
	Author    SimpleUserDTO `json:"author"`
	Reactions []ReactionDTO `json:"reactions"`
	CreatedAt time.Time     `json:"createdAt"`
}
