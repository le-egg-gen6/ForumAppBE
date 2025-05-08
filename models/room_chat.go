package models

import (
	"gorm.io/gorm"
)

type RoomChat struct {
	gorm.Model
	Name     string         `gorm:"not null"`
	Type     string         `gorm:"size:255;not null"`
	Users    []*User        `gorm:"many2many:user_room;"`
	Messages []*RoomMessage `gorm:"foreignKey:RoomID"`
}

func (*RoomChat) TableName() string {
	return "room_chat"
}
