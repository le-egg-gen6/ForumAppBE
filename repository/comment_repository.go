package repository

import (
	"forum/models"
	"gorm.io/gorm"
)

type ICommentRepository interface {
	CreateComment(comment *models.Comment) (*models.Comment, error)
	GetByID(id int64) (*models.Comment, error)
	GetByPostID(id int64) ([]models.Comment, error)
	UpdateComment(comment *models.Comment) error
	DeleteComment(id int64) error
}

type CommentRepository struct {
	db *gorm.DB
}

var CommentRepositoryInstance *CommentRepository

func InitializeCommentRepository(db *gorm.DB) {
	err := db.AutoMigrate(&models.Comment{})
	if err != nil {
		//
	}
	CommentRepositoryInstance = &CommentRepository{
		db: db,
	}
}

func GetCommentRepositoryInstance() *CommentRepository {
	return CommentRepositoryInstance
}

func (r *CommentRepository) CreateComment(comment *models.Comment) (*models.Comment, error) {
	if err := r.db.Create(comment).Error; err != nil {
		return nil, err
	}
	return comment, nil
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
