package repository

import (
	"forum/models"
	"gorm.io/gorm"
)

type IReactionRepository interface {
	Create(reaction *models.ContentReaction) (*models.ContentReaction, error)
	Update(reaction *models.ContentReaction) error
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

func (r *ReactionRepository) Create(reaction *models.ContentReaction) (*models.ContentReaction, error) {
	if err := r.db.Model(&models.ContentReaction{}).Create(reaction).Error; err != nil {
		return nil, err
	}
	return reaction, nil
}
func (r *ReactionRepository) Update(reaction *models.ContentReaction) error {
	return r.db.Model(&models.ContentReaction{}).Save(reaction).Error
}
