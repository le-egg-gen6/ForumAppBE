package repository

import (
	"gorm.io/gorm"
	"myproject/forum/models"
)

type IUserRepository interface {
	Create(user *models.User) error
	FindByID(id uint64) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	Update(user *models.User) error
	Delete(id uint64) error
	ListAll() ([]models.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		//
	}
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) FindByID(id uint64) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Update(user *models.User) error {
	return r.db.Save(user).Error
}

func (r *UserRepository) Delete(id uint64) error {
	return r.db.Model(&models.User{}).Where("id = ?", id).Update("deleted", true).Error
}

func (r *UserRepository) ListAll() ([]models.User, error) {
	var users []models.User
	if err := r.db.Where("deleted = ?", false).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
