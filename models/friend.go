package models

import (
	"gorm.io/gorm"
	"time"
)

type Friend struct {
	gorm.Model
	ID        uint64     `gorm:"primaryKey;autoIncrement"`
	UserID    *uint64    `gorm:"not null"`
	FriendID  uint64     `gorm:"not null"`
	CreatedAt *time.Time `gorm:"autoCreateTime:milli"`
	UpdatedAt *time.Time `gorm:"autoUpdateTime:milli"`
	Delete    bool       `gorm:"default:false"`
}

func (*Friend) TableName() string {
	return "friends"
}
