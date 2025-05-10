package dtos

import "time"

type MessageInfo struct {
	ID        uint          `json:"id"`
	Body      string        `json:"body"`
	Type      string        `json:"type"`
	Author    SimpleUserDTO `json:"author"`
	Reactions []ReactionDTO `json:"reactions"`
	CreatedAt time.Time     `json:"createdAt"`
}
