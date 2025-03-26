package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Instance *gorm.DB

func InitializeDatabaseConnection() {
	cfg, err := LoadDatabaseConfig()
	if err != nil {
		panic("Database configuration not found")
	}
	connectStr := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		cfg.DBHost,
		cfg.DBUsername,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBPort,
	)

	Instance, err = gorm.Open(postgres.Open(connectStr), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database" + err.Error())
	}
}

func GetDatabaseConnection() *gorm.DB {
	return Instance
}
