package models

import "time"

type Comment struct {
	ID        uint64     `gorm:"primary_key;auto_increment" json:"id"`
	UserID    uint64     `gorm:"not null" json:"user_id"`
	PostID    uint64     `gorm:"not null" json:"post_id"`
	Body      string     `gorm:"type:text;not null" json:"body"`
	CreatedAt *time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt *time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	Delete    bool       `gorm:"default:false" json:"deleted"`
}

func (Comment) TableName() string {
	return "comments"
}
