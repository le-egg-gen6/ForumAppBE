package models

import (
	"gorm.io/gorm"
)

type FriendRequest struct {
	gorm.Model
	UserID   uint64 `gorm:"not null"`
	SenderID uint64 `gorm:"not null"`
}

func (*FriendRequest) TableName() string {
	return "friend_requests"
}
