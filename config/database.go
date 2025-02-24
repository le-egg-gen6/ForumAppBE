package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(cfg *Config) *gorm.DB {
	connectStr := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		cfg.DB_HOST,
		cfg.DB_USERNAME,
		cfg.DB_PASSWORD,
		cfg.DB_NAME,
		cfg.DB_PORT)

	db, err := gorm.Open(postgres.Open(connectStr), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database" + err.Error())
	}

	return db
}
