package models

import "gorm.io/gorm"

type ContentReaction struct {
	gorm.Model
	ID        uint64  `gorm:"primaryKey;autoIncrement"`
	Type      string  `gorm:"not null"`
	PostID    *uint64 `gorm:""`
	CommentID *uint64 `gorm:""`
	MessageID *uint64 `gorm:""`
	StoryID   *uint64 `gorm:""`
	Count     int     `gorm:"default:0"`
}

func (*ContentReaction) TableName() string {
	return "content_reaction"
}
