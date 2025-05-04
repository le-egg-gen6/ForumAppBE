package repository

import (
	"errors"
	"forum/models"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Create(user *models.User) (*models.User, error)
	FindByID(id uint64) (*models.User, error)
	FindByIDWithPreloadedField(preloadField string, id uint64) (*models.User, error)
	FindByUsername(username string) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	Update(user *models.User) error
	Delete(id uint64) error
	FindAll() ([]*models.User, error)
	FindByPartialUsername(partialUsername string) ([]*models.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

var UserRepositoryInstance *UserRepository

func InitializeUserRepository(db *gorm.DB) {
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		panic("Error migrating user table: " + err.Error())
	}
	UserRepositoryInstance = &UserRepository{
		db: db,
	}
}

func GetUserRepositoryInstance() *UserRepository {
	return UserRepositoryInstance
}

func (r *UserRepository) Create(user *models.User) (*models.User, error) {
	if err := r.db.Model(&models.User{}).Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) FindByID(id uint64) (*models.User, error) {
	var user models.User
	if err := r.db.Model(&models.User{}).Where("delete = ?", false).First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindByIDWithPreloadedField(preloadField string, id uint64) (*models.User, error) {
	var user models.User
	if err := r.db.Model(&models.User{}).Preload(preloadField).Where("delete = ?", false).First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	if err := r.db.Model(&models.User{}).Where("username = ? AND delete = ?", username, false).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Model(&models.User{}).Where("email = ? AND delete = ?", email, false).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
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

func (r *UserRepository) FindAll() ([]*models.User, error) {
	var users []*models.User
	if err := r.db.Model(&models.User{}).Where("deleted = ?", false).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) FindByPartialUsername(partialUsername string) ([]*models.User, error) {
	var users []*models.User
	if err := r.db.Model(&models.User{}).Where("username LIKE ? AND delete = ?", "%"+partialUsername+"%", false).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
