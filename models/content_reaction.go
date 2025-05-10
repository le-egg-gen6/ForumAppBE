package models

import "gorm.io/gorm"

type ContentReaction struct {
	gorm.Model
	Type      string `gorm:"size:255;not null"`
	PostID    *uint  `gorm:""`
	CommentID *uint  `gorm:""`
	MessageID *uint  `gorm:""`
	StoryID   *uint  `gorm:""`
	Count     int    `gorm:"default:0"`
}

func (*ContentReaction) TableName() string {
	return "content_reaction"
}
