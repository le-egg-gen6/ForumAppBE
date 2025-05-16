package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username       string           `gorm:"size:255;not null;unique"`
	Email          string           `gorm:"size:255;not null;unique"`
	Password       string           `gorm:"size:255;not null"`
	Validated      bool             `gorm:"default:false"`
	ValidateCode   uint             `gorm:"default:0"`
	Avatar         *Image           `gorm:"foreignKey:UserID"`
	Posts          []*Post          `gorm:"foreignKey:UserID"`
	Comments       []*Comment       `gorm:"foreignKey:UserID"`
	Friends        []*Friend        `gorm:"foreignKey:UserID"`
	RoomChats      []*RoomChat      `gorm:"many2many:user_room"`
	FriendRequests []*FriendRequest `gorm:"foreignKey:UserID"`
	Notifications  []*Notification  `gorm:"foreignKey:UserID"`
}

func (*User) TableName() string {
	return "users"
}
