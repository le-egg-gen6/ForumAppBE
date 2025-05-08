package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	UserID    *uint64            `gorm:""`
	PostID    *uint64            `gorm:""`
	Body      string             `gorm:"type:text;not null"`
	Image     *Image             `gorm:"foreignKey:CommentID"`
	Reactions []*ContentReaction `gorm:"foreignKey:CommentID"`
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
