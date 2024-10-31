package db

import (
	"database/sql"
	"fmt"
	"user/domain"

	"gorm.io/gorm"
)

func MigrateUser(db *gorm.DB) error {
	// Migrate the schema
	return db.AutoMigrate(&domain.User{})
}

func SeedUser(db *gorm.DB) error {
	// Create a new user
	newUser := domain.User{
		Username: "testuser",
		Password: "password123",
		Email:    "testuser@example.com",
		Phone:    "1234567890",
	}

	// Create the user
	if err := db.Create(&newUser).Error; err != nil {
		return err
	}

	return nil
}