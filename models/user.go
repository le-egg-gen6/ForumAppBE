package models

import "time"

type User struct {
	ID           uint64     `gorm:"primary_key;auto_increment" json:"id"`
	Username     string     `gorm:"size:255;not null;unique" json:"username"`
	Email        string     `gorm:"size:255;not null;unique" json:"email"`
	Password     string     `gorm:"size:255;not null" json:"password"`
	Avatar       *Image     `gorm:"foreignKey:UserID" json:"avatar"`
	CreatedAt    *time.Time `gorm:"autoCreateTime:milli" json:"created_at"`
	UpdatedAt    *time.Time `gorm:"autoUpdateTime:milli" json:"updated_at"`
	Validated    bool       `gorm:"default:false" json:"validated"`
	ValidateCode uint64     `gorm:"default:0" json:"validate_code"`
	Deleted      bool       `gorm:"default:false" json:"deleted"`
}

func (*User) TableName() string {
	return "users"
}
