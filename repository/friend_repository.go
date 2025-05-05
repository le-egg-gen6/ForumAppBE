package repository

import (
	"forum/models"
	"gorm.io/gorm"
)

type IFriendRepository interface {
	Create(friend *models.Friend) (*models.Friend, error)
	Delete(friend *models.Friend) error
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
	if err := f.db.Model(&models.Friend{}).Create(friend).Error; err != nil {
		return nil, err
	}
	return friend, nil
}

func (f *FriendRepository) Delete(id uint64) error {
	return f.db.Model(&models.Friend{}).Where("id = ?", id).Update("delete", true).Error
}
