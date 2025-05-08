package repository

import (
	"errors"
	"forum/models"
	"gorm.io/gorm"
)

type IRoomMessageRepository interface {
	Create(roomMessage *models.RoomMessage) (*models.RoomMessage, error)
	FindByID(id uint) (*models.RoomMessage, error)
	FindByIDWithPreloadedField(id uint, preloadedField ...string) (*models.RoomMessage, error)
	Update(roomMessage *models.RoomMessage) error
	Delete(id uint) error
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
	if err := r.db.Create(roomMessage).Error; err != nil {
		return nil, err
	}
	return roomMessage, nil
}

func (r *RoomMessageRepository) FindByID(id uint) (*models.RoomMessage, error) {
	var roomMessage models.RoomMessage
	if err := r.db.First(&roomMessage, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &roomMessage, nil
}

func (r *RoomMessageRepository) FindByIDWithPreloadedField(id uint, preloadedField ...string) (*models.RoomMessage, error) {
	var roomMessage models.RoomMessage
	tx := r.db
	for _, field := range preloadedField {
		tx = tx.Preload(field)
	}
	if err := tx.First(&roomMessage, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &roomMessage, nil
}

func (r *RoomMessageRepository) Update(roomMessage *models.RoomMessage) error {
	return r.db.Model(roomMessage).Updates(roomMessage).Error
}

func (r *RoomMessageRepository) Delete(id uint) error {
	var roomMessage models.RoomMessage
	if err := r.db.First(&roomMessage, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}
	return r.db.Delete(&roomMessage).Error
}
