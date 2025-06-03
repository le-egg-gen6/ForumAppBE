package repository

import (
	"errors"
	"forum/models"
	"gorm.io/gorm"
)

type IFriendRequestRepository interface {
	Create(friendRequest *models.FriendRequest) (*models.FriendRequest, error)
	Delete(id uint) error
	IsAddedFriendRequest(userId1 uint, userId2 uint) bool
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

func GetFriendRequestRepositoryInstance() *FriendRequestRepository {
	return FriendRequestRepositoryInstance
}

func (f *FriendRequestRepository) Create(friendRequest *models.FriendRequest) (*models.FriendRequest, error) {
	if err := f.db.Create(friendRequest).Error; err != nil {
		return nil, err
	}
	return friendRequest, nil
}

func (f *FriendRequestRepository) Delete(id uint) error {
	var friendRequest models.FriendRequest
	if err := f.db.First(&friendRequest, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}
	return f.db.Delete(&friendRequest).Error
}

func (f *FriendRequestRepository) IsAddedFriendRequest(userId uint, senderId uint) bool {
	var friendRequest *models.FriendRequest = nil
	if err := f.db.Where("user_id = ? AND sender_id = ?", userId, senderId).First(friendRequest).Error; err != nil {
		return false
	}
	return true
}
