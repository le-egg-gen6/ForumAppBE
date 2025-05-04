package models

import "time"

type RoomMessage struct {
	ID        uint64             `gorm:"primary_key;auto_increment" json:"id"`
	UserID    *uint64            `gorm:"not null" json:"user_id"`
	RoomID    *string            `gorm:"not null" json:"room_id"`
	Body      string             `gorm:"type:text;not null" json:"body"`
	Image     *Image             `gorm:"foreignKey:MessageID" json:"image"`
	Reactions []*ContentReaction `gorm:"foreignKey:MessageID" json:"reactions"`
	CreatedAt *time.Time         `gorm:"autoCreateTime:milli" json:"created_at"`
	UpdatedAt *time.Time         `gorm:"autoUpdateTime:milli" json:"updated_at"`
	Delete    bool               `gorm:"default:false" json:"deleted"`
}

func (*RoomMessage) TableName() string {
	return "room_message"
}
