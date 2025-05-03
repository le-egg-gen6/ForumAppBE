package models

type ContentReaction struct {
	ID        uint64  `gorm:"primaryKey;autoIncrement" json:"id"`
	Type      string  `gorm:"not null" json:"type"`
	PostID    *uint64 `json:"post_id,omitempty"`
	CommentID *uint64 `json:"comment_id,omitempty"`
	MessageID *uint64 `json:"message_id,omitempty"`
	StoryID   *uint64 `json:"story_id,omitempty"`
	Count     int     `gorm:"default:0" json:"count"`
}

func (*ContentReaction) TableName() string {
	return "content_reaction"
}
