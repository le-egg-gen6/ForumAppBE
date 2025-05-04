package models

import (
	"gorm.io/gorm"
	"time"
)

type RoomChat struct {
	gorm.Model
	ID        uint64         `gorm:"primaryKey;autoIncrement"`
	Name      string         `gorm:"not null"`
	CreatedAt *time.Time     `gorm:"autoCreateTime:milli"`
	UpdatedAt *time.Time     `gorm:"autoUpdateTime:milli"`
	Delete    bool           `gorm:"default:false"`
	Users     []*User        `gorm:"many2many:user_room;"`
	Messages  []*RoomMessage `gorm:"foreignKey:RoomID"`
}

func (*RoomChat) TableName() string {
	return "room_chat"
}
