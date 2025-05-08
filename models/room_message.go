package models

import (
	"gorm.io/gorm"
)

type RoomMessage struct {
	gorm.Model
	UserID    *uint64            `gorm:""`
	RoomID    *uint64            `gorm:""`
	Type      string             `gorm:"size:255;not null"`
	Body      string             `gorm:"type:text;not null"`
	Reactions []*ContentReaction `gorm:"foreignKey:MessageID"`
}

func (*RoomMessage) TableName() string {
	return "room_message"
}
