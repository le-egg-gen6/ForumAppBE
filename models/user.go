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
	ValidateCode   uint64           `gorm:"default:0"`
	Posts          []*Post          `gorm:"foreignKey:UserID"`
	Comments       []*Comment       `gorm:"foreignKey:UserID"`
	Friends        []*Friend        `gorm:"foreignKey:UserID"`
	FriendRequests []*FriendRequest `gorm:"foreignKey:UserID"`
	RoomChats      []*RoomChat      `gorm:"many2many:user_room"`
}

func (*User) TableName() string {
	return "users"
}
