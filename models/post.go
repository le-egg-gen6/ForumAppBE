package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Content   string             `gorm:"type:text;not null"`
	UserID    *uint64            `gorm:""`
	Images    []*Image           `gorm:"foreignKey:PostID"`
	Comments  []*Comment         `gorm:"foreignKey:PostID"`
	Reactions []*ContentReaction `gorm:"foreignKey:PostID"`
}

func (*Post) TableName() string {
	return "posts"
}

func (p *Post) GetReaction(reactionType string) *ContentReaction {
	for _, reaction := range p.Reactions {
		if reactionType == reaction.Type {
			return reaction
		}
	}
	return nil
}
