package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func ConnectDB(cfg *Config) *gorm.DB {
	connect_str := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DB_USERNAME,
		cfg.DB_PASSWORD,
		cfg.DB_HOST,
		cfg.DB_PORT,
		cfg.DB_NAME)

	db, err := gorm.Open(postgres.Open(connect_str), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error connecting to database, %s", err)
	}

	return db
}
