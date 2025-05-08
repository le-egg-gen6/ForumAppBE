package repository

import (
	"errors"
	"forum/models"
	"gorm.io/gorm"
)

type IFriendRepository interface {
	Create(friend *models.Friend) (*models.Friend, error)
	Delete(id uint) error
}

type FriendRepository struct {
	db *gorm.DB
}

var FriendRepositoryInstance *FriendRepository

func InitializeFriendRepository(db *gorm.DB) {
	err := db.AutoMigrate(&models.Friend{})
	if err != nil {
		panic("Error migrating friend table: " + err.Error())
	}
	FriendRepositoryInstance = &FriendRepository{
		db: db,
	}
}

func GetFriendRepositoryInstance() *FriendRepository {
	return FriendRepositoryInstance
}

func (f *FriendRepository) Create(friend *models.Friend) (*models.Friend, error) {
	if err := f.db.Create(friend).Error; err != nil {
		return nil, err
	}
	return friend, nil
}

func (f *FriendRepository) Delete(id uint) error {
	var friend models.Friend
	if err := f.db.First(&friend, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}
	return f.db.Delete(&friend).Error
}
