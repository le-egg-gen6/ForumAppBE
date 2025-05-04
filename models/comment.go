package models

import "time"

type Comment struct {
	ID        uint64             `gorm:"primary_key;auto_increment" json:"id"`
	UserID    *uint64            `gorm:"not null" json:"user_id"`
	PostID    *uint64            `gorm:"not null" json:"post_id"`
	Body      string             `gorm:"type:text;not null" json:"body"`
	Image     *Image             `gorm:"foreignKey:CommentID" json:"image"`
	Reactions []*ContentReaction `gorm:"foreignKey:CommentID" json:"reactions"`
	CreatedAt *time.Time         `gorm:"autoCreateTime:milli" json:"created_at"`
	UpdatedAt *time.Time         `gorm:"autoUpdateTime:milli" json:"updated_at"`
	Delete    bool               `gorm:"default:false" json:"deleted"`
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
