package repository

import (
	"errors"
	"forum/models"
	"gorm.io/gorm"
)

type ICommentRepository interface {
	Create(comment *models.Comment) (*models.Comment, error)
	FindByID(id uint) (*models.Comment, error)
	FindByIDWithPreloadedField(id uint, preloadedField ...string) (*models.Comment, error)
	Update(comment *models.Comment) error
	Delete(id uint) error
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

func (r *CommentRepository) FindByID(id uint) (*models.Comment, error) {
	var comment models.Comment
	if err := r.db.First(&comment, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &comment, nil
}

func (r *CommentRepository) FindByIDWithPreloadedField(id uint, preloadedField ...string) (*models.Comment, error) {
	var comment models.Comment
	tx := r.db
	for _, field := range preloadedField {
		tx = tx.Preload(field)
	}
	if err := tx.First(&comment, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &comment, nil
}

func (r *CommentRepository) Update(comment *models.Comment) error {
	return r.db.Model(comment).Save(comment).Error
}

func (r *CommentRepository) Delete(id uint) error {
	var comment models.Comment
	if err := r.db.First(&comment, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}
	return r.db.Delete(&comment).Error
}
