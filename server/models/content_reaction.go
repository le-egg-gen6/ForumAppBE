package models

const (
	TypeComment = iota
	TypePost
)

type ContentReaction struct {
	ID           uint64 `gorm:"primary_key;auto_increment" json:"id"`
	ContentID    uint64 `gorm:"" json:"content_id"`
	ContentType  int    `gorm:"" json:"content_type"`
	ReactionType string `gorm:"not null" json:"reaction_type"`
	Count        int    `gorm:"" json:"count"`
}

func (*ContentReaction) TableName() string {
	return "content_reaction"
}
