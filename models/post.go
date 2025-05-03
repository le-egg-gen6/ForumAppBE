package models

import "time"

type Post struct {
	ID        uint64             `gorm:"primary_key;auto_increment" json:"id"`
	Content   string             `gorm:"type:text;not null" json:"content"`
	AuthorID  uint64             `gorm:"not null" json:"author_id"`
	Images    []*Image           `gorm:"foreignKey:PostID" json:"images"`
	Reactions []*ContentReaction `gorm:"foreignKey:PostID" json:"reactions"`
	CreatedAt *time.Time         `gorm:"autoCreateTime:milli" json:"created_at"`
	UpdatedAt *time.Time         `gorm:"autoUpdateTime:milli" json:"updated_at"`
	Deleted   bool               `gorm:"default:false" json:"deleted"`
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
