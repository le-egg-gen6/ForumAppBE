package repository

import (
	"gorm.io/gorm"
	"myproject/forum/server/models"
)

type IPostRepository interface {
	CreatePost(post *models.Post) error
	GetPostByID(id uint64) (*models.Post, error)
	GetAllPosts() ([]models.Post, error)
	UpdatePost(post *models.Post) error
	DeletePost(id uint64) error
}

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{
		db: db,
	}
}

func (r *PostRepository) CreatePost(post *models.Post) error {
	return r.db.Create(post).Error
}

func (r *PostRepository) GetPostByID(id uint64) (*models.Post, error) {
	var post models.Post
	if err := r.db.First(&post, id).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *PostRepository) GetAllPosts() ([]models.Post, error) {
	var posts []models.Post
	if err := r.db.Where("deleted = ?", false).Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *PostRepository) UpdatePost(post *models.Post) error {
	return r.db.Save(post).Error
}

func (r *PostRepository) DeletePost(id uint64) error {
	return r.db.Model(&models.Post{}).Where("id = ?", id).Update("deleted", true).Error
}
