package repository

import (
	"errors"
	"forum/models"
	"gorm.io/gorm"
)

type IRoomChatRepository interface {
	Create(roomChat *models.RoomChat) (*models.RoomChat, error)
	FindByID(id uint64) (*models.RoomChat, error)
	FindByIDWithPreloadedField(preloadField string, id uint64) (*models.RoomChat, error)
	Update(roomChat *models.RoomChat) error
	Delete(id uint64) error
}

type RoomChatRepository struct {
	db *gorm.DB
}

var RoomChatRepositoryInstance *RoomChatRepository

func InitializeRoomChatRepository(db *gorm.DB) {
	err := db.AutoMigrate(&models.RoomChat{})
	if err != nil {
		panic("Error migrating room chat table: " + err.Error())
	}
	RoomChatRepositoryInstance = &RoomChatRepository{
		db: db,
	}
}

func GetRoomChatRepositoryInstance() *RoomChatRepository {
	return RoomChatRepositoryInstance
}

func (r *RoomChatRepository) Create(roomChat *models.RoomChat) (*models.RoomChat, error) {
	if err := r.db.Model(&models.RoomChat{}).Create(roomChat).Error; err != nil {
		return nil, err
	}
	return roomChat, nil
}

func (r *RoomChatRepository) FindByID(id uint64) (*models.RoomChat, error) {
	var roomChat models.RoomChat
	if err := r.db.Model(&models.RoomChat{}).Where("delete = ?", false).First(&roomChat, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &roomChat, nil
}

func (r *RoomChatRepository) FindByIDWithPreloadedField(preloadField string, id uint64) (*models.RoomChat, error) {
	var roomChat models.RoomChat
	if err := r.db.Model(&models.RoomChat{}).Preload(preloadField).Where("delete = ?", false).First(&roomChat, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &roomChat, nil
}

func (r *RoomChatRepository) Update(roomChat *models.RoomChat) error {
	return r.db.Model(&models.RoomChat{}).Save(roomChat).Error
}

func (r *RoomChatRepository) Delete(id uint64) error {
	return r.db.Model(&models.RoomChat{}).Where("id = ?", id).Update("delete", true).Error
}
