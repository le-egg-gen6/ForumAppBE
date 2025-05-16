package repository

import (
	"errors"
	"forum/models"
	"gorm.io/gorm"
)

type INotificationRepository interface {
	Create(notification *models.Notification) (*models.Notification, error)
	Delete(id uint) error
}

type NotificationRepository struct {
	db *gorm.DB
}

var NotificationRepositoryInstance *NotificationRepository

func InitializeNotificationRepository(db *gorm.DB) {
	err := db.AutoMigrate(&models.Notification{})
	if err != nil {
		panic("Error migrating notification table: " + err.Error())
	}
	NotificationRepositoryInstance = &NotificationRepository{
		db: db,
	}
}

func GetNotificationRepositoryInstance() *NotificationRepository {
	return NotificationRepositoryInstance
}

func (r *NotificationRepository) Create(notification *models.Notification) (*models.Notification, error) {
	if err := r.db.Create(notification).Error; err != nil {
		return nil, err
	}
	return notification, nil
}

func (r *NotificationRepository) Delete(id uint) error {
	var notification models.Notification
	if err := r.db.First(&notification, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}
	return r.db.Delete(&notification).Error
}
