package dtos

import "time"

type CommentDTO struct {
	ID        uint64        `json:"id"`
	Author    SimpleUserDTO `json:"author"`
	PostID    uint64        `json:"postId"`
	Body      string        `json:"body"`
	CreatedAt time.Time     `json:"createdAt"`
	Reactions []ReactionDTO `json:"reactions"`
}

type SimpleCommentDTO struct {
	ID     uint64        `json:"id"`
	Author SimpleUserDTO `json:"author"`
	PostID uint64        `json:"postId"`
	Body   string        `json:"body"`
}
