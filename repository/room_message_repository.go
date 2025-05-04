package repository

import (
	"errors"
	"forum/models"
	"gorm.io/gorm"
)

type IRoomMessageRepository interface {
	Create(roomMessage *models.RoomMessage) (*models.RoomMessage, error)
	FindByID(id int64) (*models.RoomMessage, error)
	FindByIDWithPreloadedField(preloadField string, id uint64) (*models.RoomMessage, error)
	Update(roomMessage *models.RoomMessage) error
	Delete(id int64) error
}

type RoomMessageRepository struct {
	db *gorm.DB
}

var RoomMessageRepositoryInstance *RoomMessageRepository

func InitializeRoomMessageRepository(db *gorm.DB) {
	err := db.AutoMigrate(&models.RoomMessage{})
	if err != nil {
		panic("Error migrating room message table: " + err.Error())
	}
	RoomMessageRepositoryInstance = &RoomMessageRepository{
		db: db,
	}
}

func GetRoomMessageRepositoryInstance() *RoomMessageRepository {
	return RoomMessageRepositoryInstance
}

func (r *RoomMessageRepository) Create(roomMessage *models.RoomMessage) (*models.RoomMessage, error) {
	if err := r.db.Model(&models.RoomMessage{}).Create(roomMessage).Error; err != nil {
		return nil, err
	}
	return roomMessage, nil
}

func (r *RoomMessageRepository) FindByID(id int64) (*models.RoomMessage, error) {
	var roomMessage models.RoomMessage
	if err := r.db.Model(&models.RoomMessage{}).Where("delete = ?", false).First(&roomMessage, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &roomMessage, nil
}

func (r *RoomMessageRepository) FindByIDWithPreloadedField(preloadField string, id uint64) (*models.RoomMessage, error) {
	var roomMessage models.RoomMessage
	if err := r.db.Model(&models.RoomMessage{}).Preload(preloadField).Where("delete = ?", false).First(&roomMessage, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &roomMessage, nil
}

func (r *RoomMessageRepository) Update(roomMessage *models.RoomMessage) error {
	return r.db.Model(&models.RoomMessage{}).Save(roomMessage).Error
}

func (r *RoomMessageRepository) Delete(id int64) error {
	return r.db.Model(&models.RoomMessage{}).Where("id = ?", id).Update("delete", true).Error
}
