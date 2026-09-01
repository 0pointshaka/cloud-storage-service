package config

import (
	"os"

	"github.com/0pointshaka/cloud-storage-service/backend/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	db, err := gorm.Open(sqlite.Open("cloud_storage.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	// Auto-migrate models
	if err := db.AutoMigrate(
		&models.User{},
		&models.File{},
		&models.SharedFile{},
	); err != nil {
		return err
	}

	DB = db
	return nil
}

func LoadEnv() error {
	if os.Getenv("JWT_SECRET") == "" {
		os.Setenv("JWT_SECRET", "your-secret-key-change-in-production")
	}
	return nil
}
