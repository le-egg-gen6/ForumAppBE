package repository

import (
	"errors"
	"forum/models"
	"gorm.io/gorm"
)

type IHashTagRepository interface {
	FindAll(limit int) ([]*models.HashTag, error)
	IncrementCount(name string) error
}

type HashTagRepository struct {
	db *gorm.DB
}

var HashTagRepositoryInstance *HashTagRepository

func InitializeHashTagRepository(db *gorm.DB) {
	err := db.AutoMigrate(&models.HashTag{})
	if err != nil {
		panic("Error migrating hash tag table: " + err.Error())
	}
	HashTagRepositoryInstance = &HashTagRepository{
		db: db,
	}
}

func GetHashTagRepositoryInstance() *HashTagRepository {
	return HashTagRepositoryInstance
}

func (r *HashTagRepository) FindAll(limit int) ([]*models.HashTag, error) {
	var hashTags []*models.HashTag
	err := r.db.Order("count desc").Limit(limit).Find(&hashTags).Error
	return hashTags, err
}

func (r *HashTagRepository) IncrementCount(name string) error {
	var hashTag models.HashTag
	if err := r.db.First(&hashTag, "name = ?", name).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			hashTag = models.HashTag{
				Name:  name,
				Count: 1,
			}
			return r.db.Create(&hashTag).Error
		}
		return err
	}
	hashTag.Count++
	return r.db.Save(&hashTag).Error
}
