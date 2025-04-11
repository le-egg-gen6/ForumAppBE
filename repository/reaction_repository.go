package repository

import (
	"forum/models"
	"gorm.io/gorm"
)

type IReactionRepository interface {
	FindByContentIDAndContentType(contentID uint64, contentType int) ([]models.ContentReaction, error)
}

type ReactionRepository struct {
	db *gorm.DB
}

var ReactionRepositoryInstance *ReactionRepository

func InitializeReactionRepository(db *gorm.DB) {
	err := db.AutoMigrate(&models.ContentReaction{})
	if err != nil {
		panic("Error migrating reaction table: " + err.Error())
	}
	ReactionRepositoryInstance = &ReactionRepository{
		db: db,
	}
}

func GetReactionRepositoryInstance() *ReactionRepository {
	return ReactionRepositoryInstance
}

func (r *ReactionRepository) FindByContentIDAndContentType(contentId uint64, contentType int) ([]models.ContentReaction, error) {
	var reactions []models.ContentReaction
	if err := r.db.Where("content_id = ? AND content_type = ?", contentId, contentType).Find(&reactions).Error; err != nil {
		return nil, err
	}
	return reactions, nil
}
