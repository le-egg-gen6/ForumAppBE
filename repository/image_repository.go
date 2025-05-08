package repository

import (
	"errors"
	"forum/models"
	"gorm.io/gorm"
)

type IImageRepository interface {
	Create(image *models.Image) (*models.Image, error)
	FindByID(id uint) (*models.Image, error)
}

type ImageRepository struct {
	db *gorm.DB
}

var ImageRepositoryInstance *ImageRepository

func InitializeImageRepository(db *gorm.DB) {
	err := db.AutoMigrate(&models.Image{})
	if err != nil {
		panic("Error migrating image table: " + err.Error())
	}
	ImageRepositoryInstance = &ImageRepository{
		db: db,
	}
}

func GetImageRepositoryInstance() *ImageRepository {
	return ImageRepositoryInstance
}

func (r *ImageRepository) Create(image *models.Image) (*models.Image, error) {
	if err := r.db.Create(image).Error; err != nil {
		return nil, err
	}
	return image, nil
}

func (r *ImageRepository) FindByID(id uint) (*models.Image, error) {
	var image models.Image
	if err := r.db.First(&image, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &image, nil
}
