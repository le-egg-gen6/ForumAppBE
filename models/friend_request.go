package models

import (
	"gorm.io/gorm"
	"time"
)

type FriendRequest struct {
	gorm.Model
	ID        uint64     `gorm:"primaryKey;autoIncrement"`
	UserID    *uint64    `gorm:"not null"`
	SenderID  uint64     `gorm:"not null"`
	CreatedAt *time.Time `gorm:"autoCreateTime:milli"`
	UpdatedAt *time.Time `gorm:"autoUpdateTime:milli"`
	Delete    bool       `gorm:"default:false"`
}

func (*FriendRequest) TableName() string {
	return "friend_requests"
}
