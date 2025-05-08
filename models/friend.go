package models

import (
	"gorm.io/gorm"
)

type Friend struct {
	gorm.Model
	UserID   uint64 `gorm:"not null"`
	FriendID uint64 `gorm:"not null"`
}

func (*Friend) TableName() string {
	return "friends"
}
