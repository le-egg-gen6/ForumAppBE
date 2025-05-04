package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	ID           uint64      `gorm:"primaryKey;autoIncrement"`
	Username     string      `gorm:"size:255;not null;unique"`
	Email        string      `gorm:"size:255;not null;unique"`
	Password     string      `gorm:"size:255;not null"`
	Avatar       *Image      `gorm:"foreignKey:UserID"`
	CreatedAt    *time.Time  `gorm:"autoCreateTime:milli"`
	UpdatedAt    *time.Time  `gorm:"autoUpdateTime:milli"`
	Validated    bool        `gorm:"default:false"`
	ValidateCode uint64      `gorm:"default:0"`
	Deleted      bool        `gorm:"default:false"`
	Posts        []*Post     `gorm:"foreignKey:AuthorID"`
	Comments     []*Comment  `gorm:"foreignKey:UserID"`
	RoomChats    []*RoomChat `gorm:"many2many:user_room"`
}

func (*User) TableName() string {
	return "users"
}
