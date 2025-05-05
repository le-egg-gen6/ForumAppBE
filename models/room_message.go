package models

import (
	"gorm.io/gorm"
	"time"
)

type RoomMessage struct {
	gorm.Model
	ID        uint64             `gorm:"primaryKey;autoIncrement"`
	UserID    *uint64            `gorm:"not null"`
	RoomID    *string            `gorm:"not null"`
	Body      string             `gorm:"type:text;not null"`
	Reactions []*ContentReaction `gorm:"foreignKey:MessageID"`
	CreatedAt *time.Time         `gorm:"autoCreateTime:milli"`
	UpdatedAt *time.Time         `gorm:"autoUpdateTime:milli"`
	Delete    bool               `gorm:"default:false"`
}

func (*RoomMessage) TableName() string {
	return "room_message"
}
