package models

import "time"

type LikeComment struct {
	ID        uint64     `gorm:"primary_key;auto_increment" json:"id"`
	UserID    uint64     `gorm:"not null" json:"user_id"`
	CommentID uint64     `gorm:"not null" json:"comment_id"`
	CreatedAt *time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt *time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	Deleted   bool       `gorm:"default:false" json:"deleted"`
}

func (*LikeComment) TableName() string {
	return "like_comments"
}
