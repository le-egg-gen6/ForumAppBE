package models

import "time"

type RoomChat struct {
	ID        string         `gorm:"primary_key;not null" json:"id"`
	Name      string         `gorm:"not null" json:"name"`
	CreatedAt *time.Time     `gorm:"autoCreateTime:milli" json:"created_at"`
	UpdatedAt *time.Time     `gorm:"autoUpdateTime:milli" json:"updated_at"`
	Delete    bool           `gorm:"default:false" json:"deleted"`
	Messages  []*RoomMessage `gorm:"foreignKey:RoomID" json:"messages"`
}

func (*RoomChat) TableName() string {
	return "room_chat"
}
