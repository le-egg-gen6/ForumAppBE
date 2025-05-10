package repository

import (
	"errors"
	"forum/models"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Create(user *models.User) (*models.User, error)
	FindByID(id uint) (*models.User, error)
	FindByIDWithPreloadedField(id uint, preloadField ...string) (*models.User, error)
	FindByUsername(username string) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	Update(user *models.User) error
	UpdateAssociations(user *models.User, associationField string, obj ...interface{}) error
	Delete(id uint) error
	DeleteAssociations(user *models.User, associationField string, obj ...interface{}) error
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
	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) FindByID(id uint) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindByIDWithPreloadedField(id uint, preloadField ...string) (*models.User, error) {
	var user models.User
	tx := r.db
	for _, field := range preloadField {
		tx = tx.Preload(field)
	}
	if err := tx.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Update(user *models.User) error {
	return r.db.Model(user).Updates(user).Error
}

func (r *UserRepository) UpdateAssociations(user *models.User, associationField string, objs ...interface{}) error {
	return r.db.Model(user).Association(associationField).Append(objs)
}

func (r *UserRepository) Delete(id uint) error {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}
	return r.db.Delete(&user).Error
}

func (r *UserRepository) DeleteAssociations(user *models.User, associationField string, objs ...interface{}) error {
	return r.db.Model(user).Association(associationField).Delete(objs)
}

func (r *UserRepository) FindAll() ([]*models.User, error) {
	var users []*models.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) FindByPartialUsername(partialUsername string) ([]*models.User, error) {
	var users []*models.User
	if err := r.db.Where("username LIKE ?", "%"+partialUsername+"%").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
