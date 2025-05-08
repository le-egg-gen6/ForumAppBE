package repository

import (
	"errors"
	"forum/models"
	"gorm.io/gorm"
)

type IPostRepository interface {
	Create(post *models.Post) (*models.Post, error)
	FindByID(id uint) (*models.Post, error)
	FindByIDWithPreloadedField(id uint, preloadedField ...string) (*models.Post, error)
	FindAll() ([]*models.Post, error)
	Update(post *models.Post) error
	Delete(id uint) error
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
	if err := r.db.Create(post).Error; err != nil {
		return nil, err
	}
	return post, nil
}

func (r *PostRepository) FindByID(id uint) (*models.Post, error) {
	var post models.Post
	if err := r.db.First(&post, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &post, nil
}

func (r *PostRepository) FindByIDWithPreloadedField(id uint, preloadedField ...string) (*models.Post, error) {
	var post models.Post
	tx := r.db
	for _, field := range preloadedField {
		tx.Preload(field)
	}
	if err := tx.First(&post, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &post, nil
}

func (r *PostRepository) FindAll() ([]*models.Post, error) {
	var posts []*models.Post
	if err := r.db.Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *PostRepository) Update(post *models.Post) error {
	return r.db.Model(post).Updates(post).Error
}

func (r *PostRepository) Delete(id uint) error {
	var post models.Post
	if err := r.db.First(&post, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}
	return r.db.Delete(&post).Error
}
