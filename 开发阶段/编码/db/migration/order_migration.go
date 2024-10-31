package db

import (
	"database/sql"
	"fmt"
	"order/domain"

	"gorm.io/gorm"
)

func MigrateOrder(db *gorm.DB) error {
	// Migrate the schema
	return db.AutoMigrate(&domain.Order{})
}

func SeedOrder(db *gorm.DB) error {
	// Create a new order
	newOrder := domain.Order{
		UserID:   1,
		GoodsID:  1,
		Quantity: 2,
		Price:    200.00,
		Status:   1, // 1: Pending
	}

	// Create the order
	if err := db.Create(&newOrder).Error; err != nil {
		return err
	}

	return nil
}