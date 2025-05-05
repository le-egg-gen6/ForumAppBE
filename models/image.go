package models

import (
	"gorm.io/gorm"
	"time"
)

type Image struct {
	gorm.Model
	ID        uint64     `gorm:"primaryKey;autoIncrement"`
	URL       string     `gorm:"type:text;not null"`
	UserID    *uint64    `gorm:""`
	PostID    *uint64    `gorm:""`
	CommentID *uint64    `gorm:""`
	CreatedAt *time.Time `gorm:"autoCreateTime:milli"`
}

func (*Image) TableName() string {
	return "image"
}
