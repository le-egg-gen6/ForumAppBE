package models

import "time"

type User struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Username  string    `gorm:"size:255;not null;unique" json:"username"`
	Email     string    `gorm:"size:255;not null;unique" json:"email"`
	Password  string    `gorm:"size:255;not null" json:"password"`
	AvatarUrl string    `gorm:"size:255;null" json:"avatar_url"`
	CreatedAt time.Time `gorm:"autoCreateTime:milli" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:milli" json:"updated_at"`
	Deleted   bool      `gorm:"default:false" json:"deleted"`
}

func (User) TableName() string {
	return "users"
}
