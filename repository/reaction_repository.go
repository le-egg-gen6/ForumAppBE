package repository

import (
	"errors"
	"gorm.io/gorm"
	"myproject/forum/models"
)

type IReactionRepository interface {
	FindByType(contentID uint64, contentType int) ([]models.ContentReaction, error)
	IncreaseReaction(contentId uint64, contentType int, reactionType string) (*models.ContentReaction, error)
	DecreaseReaction(contentId uint64, contentType int, reactionType string) (*models.ContentReaction, error)
}

type ReactionRepository struct {
	db *gorm.DB
}

func NewReactionRepository(db *gorm.DB) *ReactionRepository {
	err := db.AutoMigrate(&models.ContentReaction{})
	if err != nil {
		//
	}
	return &ReactionRepository{
		db: db,
	}
}

func (r *ReactionRepository) FindByType(contentId uint64, contentType int) ([]models.ContentReaction, error) {
	var reactions []models.ContentReaction
	if err := r.db.Where("content_id = ? AND content_type = ?", contentId, contentType).Find(&reactions).Error; err != nil {
		return nil, err
	}
	return reactions, nil
}

func (r *ReactionRepository) IncreaseReaction(contentId uint64, contentType int, reactionType string) (*models.ContentReaction, error) {
	var reaction models.ContentReaction
	err := r.db.Where("content_id = ? AND content_type = ? AND reaction_type = ?").Find(&reaction).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			reaction = models.ContentReaction{
				ContentID:    contentId,
				ContentType:  contentType,
				ReactionType: reactionType,
				Count:        1,
			}
			if err := r.db.Create(&reaction).Error; err != nil {
				return nil, err
			}
			return &reaction, nil
		}
		return nil, err
	}
	reaction.Count++
	if err := r.db.Save(&reaction).Error; err != nil {
		return nil, err
	}

	return &reaction, nil
}

func (r *ReactionRepository) DecreaseReaction(contentId uint64, contentType int, reactionType string) (*models.ContentReaction, error) {
	var reaction models.ContentReaction
	err := r.db.Where("content_id = ? AND content_type = ? AND reaction_type = ?", contentId, contentType, reactionType).
		First(&reaction).Error

	if err != nil {
		return nil, err
	}

	reaction.Count--
	if err := r.db.Save(&reaction).Error; err != nil {
		return nil, err
	}
	return &reaction, nil
}
