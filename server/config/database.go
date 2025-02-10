package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(cfg *Config) *gorm.DB {
	connectStr := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DB_USERNAME,
		cfg.DB_PASSWORD,
		cfg.DB_HOST,
		cfg.DB_PORT,
		cfg.DB_NAME)

	db, err := gorm.Open(postgres.Open(connectStr), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database" + err.Error())
	}

	return db
}
