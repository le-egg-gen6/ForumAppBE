package models

import (
	"gorm.io/gorm"
	"time"
)

type Post struct {
	gorm.Model
	ID        uint64             `gorm:"primaryKey;autoIncrement"`
	Content   string             `gorm:"type:text;not null"`
	AuthorID  *uint64            `gorm:"not null"`
	Images    []*Image           `gorm:"foreignKey:PostID"`
	Comments  []*Comment         `gorm:"foreignKey:PostID"`
	Reactions []*ContentReaction `gorm:"foreignKey:PostID"`
	CreatedAt *time.Time         `gorm:"autoCreateTime:milli"`
	UpdatedAt *time.Time         `gorm:"autoUpdateTime:milli"`
	Deleted   bool               `gorm:"default:false"`
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
