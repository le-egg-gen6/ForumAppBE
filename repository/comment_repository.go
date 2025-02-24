package repository

import (
	"gorm.io/gorm"
	"myproject/forum/models"
)

type ICommentRepository interface {
	CreateComment(comment *models.Comment) error
	GetByID(id int64) (*models.Comment, error)
	GetByPostID(id int64) ([]models.Comment, error)
	UpdateComment(comment *models.Comment) error
	DeleteComment(id int64) error
}

type CommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentRepository {
	err := db.AutoMigrate(&models.Comment{})
	if err != nil {
		//
	}
	return &CommentRepository{
		db: db,
	}
}

func (r *CommentRepository) CreateComment(comment *models.Comment) error {
	return r.db.Create(comment).Error
}

func (r *CommentRepository) GetByID(id int64) (*models.Comment, error) {
	var comment models.Comment
	if err := r.db.First(&comment, id).Error; err != nil {
		return nil, err
	}
	return &comment, nil
}

func (r *CommentRepository) GetByPostID(id int64) ([]models.Comment, error) {
	var comments []models.Comment
	if err := r.db.Where("post_id = ? AND delete = ?", id, false).Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

func (r *CommentRepository) UpdateComment(comment *models.Comment) error {
	return r.db.Save(comment).Error
}

func (r *CommentRepository) DeleteComment(id int64) error {
	return r.db.Model(&models.Comment{}).Where("id = ?", id).Update("delete", true).Error
}
