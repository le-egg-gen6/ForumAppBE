package repository

import (
	"forum/models"
	"gorm.io/gorm"
)

type ICommentRepository interface {
	Create(comment *models.Comment) (*models.Comment, error)
	FindByID(id uint64) (*models.Comment, error)
	FindByPostID(id uint64) ([]models.Comment, error)
	Update(comment *models.Comment) error
	Delete(id uint64) error
}

type CommentRepository struct {
	db *gorm.DB
}

var CommentRepositoryInstance *CommentRepository

func InitializeCommentRepository(db *gorm.DB) {
	err := db.AutoMigrate(&models.Comment{})
	if err != nil {
		panic("Error migrating comment table: " + err.Error())
	}
	CommentRepositoryInstance = &CommentRepository{
		db: db,
	}
}

func GetCommentRepositoryInstance() *CommentRepository {
	return CommentRepositoryInstance
}

func (r *CommentRepository) Create(comment *models.Comment) (*models.Comment, error) {
	if err := r.db.Create(comment).Error; err != nil {
		return nil, err
	}
	return comment, nil
}

func (r *CommentRepository) FindByID(id uint64) (*models.Comment, error) {
	var comment models.Comment
	if err := r.db.First(&comment, id).Error; err != nil {
		return nil, err
	}
	return &comment, nil
}

func (r *CommentRepository) FindByPostID(id uint64) ([]models.Comment, error) {
	var comments []models.Comment
	if err := r.db.Where("post_id = ? AND delete = ?", id, false).Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

func (r *CommentRepository) Update(comment *models.Comment) error {
	return r.db.Save(comment).Error
}

func (r *CommentRepository) Delete(id uint64) error {
	return r.db.Model(&models.Comment{}).Where("id = ?", id).Update("delete", true).Error
}
