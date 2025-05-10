package dtos

import (
	"time"
)

type PostDTO struct {
	ID          uint          `json:"id"`
	Content     string        `json:"content"`
	Author      SimpleUserDTO `json:"author"`
	Images      []string      `json:"images"`
	CreatedAt   time.Time     `json:"createdAt"`
	Reactions   []ReactionDTO `json:"reactions"`
	TopComments []CommentDTO  `json:"topComments"`
}

type SimplePostDTO struct {
	ID        uint          `json:"id"`
	Content   string        `json:"content"`
	Author    SimpleUserDTO `json:"author"`
	Images    []string      `json:"images"`
	CreatedAt time.Time     `json:"createdAt"`
}
