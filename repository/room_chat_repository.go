package repository

import (
	"errors"
	"forum/models"
	"gorm.io/gorm"
)

type IRoomChatRepository interface {
	Create(roomChat *models.RoomChat) (*models.RoomChat, error)
	FindByID(id uint) (*models.RoomChat, error)
	FindByIDWithPreloadedField(id uint, preloadedField ...string) (*models.RoomChat, error)
	Update(roomChat *models.RoomChat) error
	UpdateAssociations(roomChat *models.RoomChat, associationField string, objs ...interface{}) error
	Delete(id uint) error
	DeleteAssociations(roomChat *models.RoomChat, associationField string, objs ...interface{}) error
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
	if err := r.db.Create(roomChat).Error; err != nil {
		return nil, err
	}
	return roomChat, nil
}

func (r *RoomChatRepository) FindByID(id uint) (*models.RoomChat, error) {
	var roomChat models.RoomChat
	if err := r.db.First(&roomChat, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &roomChat, nil
}

func (r *RoomChatRepository) FindByIDWithPreloadedField(id uint, preloadedField ...string) (*models.RoomChat, error) {
	var roomChat models.RoomChat
	tx := r.db
	for _, preloadField := range preloadedField {
		tx = tx.Preload(preloadField)
	}
	if err := tx.First(&roomChat, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &roomChat, nil
}

func (r *RoomChatRepository) Update(roomChat *models.RoomChat) error {
	return r.db.Model(roomChat).Updates(roomChat).Error
}

func (r *RoomChatRepository) UpdateAssociations(roomChat *models.RoomChat, associationField string, objs ...interface{}) error {
	return r.db.Model(roomChat).Association(associationField).Append(objs)
}

func (r *RoomChatRepository) Delete(id uint) error {
	var roomChat models.RoomChat
	if err := r.db.First(&roomChat, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}
	return r.db.Delete(&roomChat).Error
}

func (r *RoomChatRepository) DeleteAssociations(roomChat *models.RoomChat, associationField string, objs ...interface{}) error {
	return r.db.Model(roomChat).Association(associationField).Delete(objs)
}
