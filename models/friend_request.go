package models

import (
	"gorm.io/gorm"
)

type FriendRequest struct {
	gorm.Model
	UserID   uint `gorm:"not null"`
	SenderID uint `gorm:"not null"`
}

func (*FriendRequest) TableName() string {
	return "friend_requests"
}
