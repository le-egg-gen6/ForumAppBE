package models

import "gorm.io/gorm"

type Notification struct {
	gorm.Model
	Content string `gorm:"type:text;not null"`
	UserID  *uint  `gorm:""`
}

func (*Notification) TableName() string {
	return "notification"
}
