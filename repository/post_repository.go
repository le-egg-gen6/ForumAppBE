package repository

import (
	"forum/models"
	"gorm.io/gorm"
)

type IPostRepository interface {
	Create(post *models.Post) (*models.Post, error)
	FindByPostID(id uint64) (*models.Post, error)
	FindAll() ([]models.Post, error)
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
	if err := r.db.Create(post).Error; err != nil {
		return nil, err
	}
	return post, nil
}

func (r *PostRepository) FindByPostID(id uint64) (*models.Post, error) {
	var post models.Post
	if err := r.db.First(&post, id).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *PostRepository) FindAll() ([]models.Post, error) {
	var posts []models.Post
	if err := r.db.Where("deleted = ?", false).Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *PostRepository) Update(post *models.Post) error {
	return r.db.Save(post).Error
}

func (r *PostRepository) Delete(id uint64) error {
	return r.db.Model(&models.Post{}).Where("id = ?", id).Update("deleted", true).Error
}
