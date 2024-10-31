package db

import (
	"database/sql"
	"fmt"
	"customer_service/domain"

	"gorm.io/gorm"
)

func MigrateCustomerService(db *gorm.DB) error {
	// Migrate the schema
	return db.AutoMigrate(&domain.CustomerFeedback{})
}

func SeedCustomerService(db *gorm.DB) error {
	// Create a new customer feedback
	newFeedback := domain.CustomerFeedback{
		UserId:    1,
		Message:   "Great product!",
		FeedbackType: "positive",
	}

	// Create the customer feedback
	if err := db.Create(&newFeedback).Error; err != nil {
		return err
	}

	return nil
}