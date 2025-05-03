package models

import "time"

type Image struct {
	ID        uint64     `gorm:"primary_key;auto_increment" json:"id"`
	URL       string     `gorm:"type:text;not null" json:"url"`
	UserID    *uint64    `json:"user_id,omitempty"`
	PostID    *uint64    `json:"post_id,omitempty"`
	CommentID *uint64    `json:"comment_id,omitempty"`
	MessageID *uint64    `json:"message_id,omitempty"`
	StoryID   *uint64    `json:"story_id,omitempty"`
	CreatedAt *time.Time `gorm:"autoCreateTime:milli" json:"created_at"`
	UpdatedAt *time.Time `gorm:"autoUpdateTime:milli" json:"updated_at"`
	Deleted   bool       `gorm:"default:false" json:"deleted"`
}

func (*Image) TableName() string {
	return "image"
}
