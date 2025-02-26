package models

type ContentType int

const (
	TypeComment ContentType = 0
	TypePost    ContentType = 1
)

type ReactionType string

const (
	Like  ReactionType = "like"
	Love  ReactionType = "love"
	Haha  ReactionType = "haha"
	Wow   ReactionType = "wow"
	Sad   ReactionType = "sad"
	Angry ReactionType = "angry"
)

type ContentReaction struct {
	ID           uint64       `gorm:"primaryKey;autoIncrement" json:"id"`
	ContentID    uint64       `gorm:"not null" json:"content_id"`
	ContentType  ContentType  `gorm:"not null" json:"content_type"`
	ReactionType ReactionType `gorm:"not null" json:"reaction_type"`
	Count        int          `gorm:"default:0" json:"count"`
}

func (ContentReaction) TableName() string {
	return "content_reaction"
}
