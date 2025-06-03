package repository

import (
	"errors"
	"forum/models"
	"gorm.io/gorm"
)

type IFriendRepository interface {
	Create(friend *models.Friend) (*models.Friend, error)
	Delete(id uint) error
	IsFriend(userId1 uint, userId2 uint) bool
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

func (f *FriendRepository) IsFriend(userId1 uint, userId2 uint) bool {
	var friend_1 *models.Friend = nil
	var friend_2 *models.Friend = nil
	err1 := f.db.Where("user_id = ? AND friend_id = ?", userId1, userId2).First(friend_1).Error
	err2 := f.db.Where("user_id = ? AND friend_id = ?", userId2, userId1).First(friend_2).Error
	if err1 != nil && err2 != nil {
		return false
	}
	return true
}
