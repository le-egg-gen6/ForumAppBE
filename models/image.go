package models

import (
	"gorm.io/gorm"
)

type Image struct {
	gorm.Model
	URL       string  `gorm:"type:text;not null"`
	UserID    *uint64 `gorm:""`
	PostID    *uint64 `gorm:""`
	CommentID *uint64 `gorm:""`
}

func (*Image) TableName() string {
	return "image"
}
