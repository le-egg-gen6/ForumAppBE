package repository

import (
	"errors"
	"forum/models"
	"gorm.io/gorm"
)

type IPostRepository interface {
	Create(post *models.Post) (*models.Post, error)
	FindByID(id uint64) (*models.Post, error)
	FindByIDWithPreloadedField(preloadField string, id uint64) (*models.Post, error)
	FindAll() ([]*models.Post, error)
	Update(post *models.Post) error
	Delete(id uint64) error
}

type PostRepository struct {
	db *gorm.DB
}

var PostRepositoryInstance *PostRepository

func InitializePostRepository(db *gorm.DB) {
	err := db.AutoMigrate(&models.Post{})
	if err != nil {
		panic("Error migrating post table: " + err.Error())
	}
	PostRepositoryInstance = &PostRepository{
		db: db,
	}
}

func GetPostRepositoryInstance() *PostRepository {
	return PostRepositoryInstance
}

func (r *PostRepository) Create(post *models.Post) (*models.Post, error) {
	if err := r.db.Model(&models.Post{}).Create(post).Error; err != nil {
		return nil, err
	}
	return post, nil
}

func (r *PostRepository) FindByID(id uint64) (*models.Post, error) {
	var post models.Post
	if err := r.db.Model(&models.Post{}).Where("delete = ?", false).First(&post, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &post, nil
}

func (r *PostRepository) FindByIDWithPreloadedField(preloadField string, id uint64) (*models.Post, error) {
	var post models.Post
	if err := r.db.Model(&models.Post{}).Preload(preloadField).Where("delete = ?", false).First(&post, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &post, nil
}

func (r *PostRepository) FindAll() ([]*models.Post, error) {
	var posts []*models.Post
	if err := r.db.Model(&models.Post{}).Where("deleted = ?", false).Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *PostRepository) Update(post *models.Post) error {
	return r.db.Model(&models.Post{}).Save(post).Error
}

func (r *PostRepository) Delete(id uint64) error {
	return r.db.Model(&models.Post{}).Where("id = ?", id).Update("deleted", true).Error
}
