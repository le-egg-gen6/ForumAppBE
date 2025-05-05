package repository

import (
	"forum/models"
	"gorm.io/gorm"
)

type IFriendRequestRepository interface {
	Create(friendRequest *models.FriendRequest) (*models.FriendRequest, error)
	Delete(id uint64) error
}

type FriendRequestRepository struct {
	db *gorm.DB
}

var FriendRequestRepositoryInstance *FriendRequestRepository

func InitializeFriendRequestRepository(db *gorm.DB) {
	err := db.AutoMigrate(&models.FriendRequest{})
	if err != nil {
		panic("Error migrating friend request table: " + err.Error())
	}
	FriendRepositoryInstance = &FriendRepository{
		db: db,
	}
}

func (f *FriendRequestRepository) Create(friendRequest *models.FriendRequest) (*models.FriendRequest, error) {
	if err := f.db.Model(&models.FriendRequest{}).Create(friendRequest).Error; err != nil {
		return nil, err
	}
	return friendRequest, nil
}

func (f *FriendRequestRepository) Delete(id uint64) error {
	return f.db.Model(&models.FriendRequest{}).Where("id = ?", id).Update("delete", true).Error
}
