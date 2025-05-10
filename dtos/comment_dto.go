package dtos

import "time"

type CommentDTO struct {
	ID        uint          `json:"id"`
	Author    SimpleUserDTO `json:"author"`
	PostID    uint          `json:"postId"`
	Body      string        `json:"body"`
	CreatedAt time.Time     `json:"createdAt"`
	Reactions []ReactionDTO `json:"reactions"`
}

type SimpleCommentDTO struct {
	ID        uint          `json:"id"`
	Author    SimpleUserDTO `json:"author"`
	PostID    uint          `json:"postId"`
	Body      string        `json:"body"`
	CreatedAt time.Time     `json:"createdAt"`
}
