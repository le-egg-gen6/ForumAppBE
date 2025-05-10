package models

import (
	"gorm.io/gorm"
)

type Image struct {
	gorm.Model
	URL       string `gorm:"type:text;not null"`
	UserID    *uint  `gorm:""`
	PostID    *uint  `gorm:""`
	CommentID *uint  `gorm:""`
}

func (*Image) TableName() string {
	return "image"
}
