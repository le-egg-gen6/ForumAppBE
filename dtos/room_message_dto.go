package dtos

import "time"

type MessageInfo struct {
	MessageID string        `json:"messageID"`
	Body      string        `json:"body"`
	Reactions []ReactionDTO `json:"reactions"`
	CreatedAt *time.Time    `json:"createdAt"`
}
