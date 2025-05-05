package models

import (
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	gorm.Model
	ID        uint64             `gorm:"primaryKey;autoIncrement"`
	UserID    *uint64            `gorm:"not null"`
	PostID    *uint64            `gorm:"not null"`
	Body      string             `gorm:"type:text;not null"`
	Image     *Image             `gorm:"foreignKey:CommentID"`
	Reactions []*ContentReaction `gorm:"foreignKey:CommentID"`
	CreatedAt *time.Time         `gorm:"autoCreateTime:milli"`
	UpdatedAt *time.Time         `gorm:"autoUpdateTime:milli"`
	Delete    bool               `gorm:"default:false"`
}

func (*Comment) TableName() string {
	return "comments"
}

func (c *Comment) GetReaction(reactionType string) *ContentReaction {
	for _, reaction := range c.Reactions {
		if reactionType == reaction.Type {
			return reaction
		}
	}
	return nil
}
