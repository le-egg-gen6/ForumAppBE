package models

import (
	"gorm.io/gorm"
)

type Friend struct {
	gorm.Model
	UserID   uint `gorm:"not null"`
	FriendID uint `gorm:"not null"`
}

func (*Friend) TableName() string {
	return "friends"
}
