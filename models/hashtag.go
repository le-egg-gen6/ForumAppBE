package models

import "gorm.io/gorm"

type HashTag struct {
	gorm.Model
	Name  string `gorm:"size:255;not null"`
	Count uint   `gorm:"default:0"`
}

func (*HashTag) TableName() string {
	return "hash_tag"
}
