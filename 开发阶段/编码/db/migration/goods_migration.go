package db

import (
	"database/sql"
	"fmt"
	"goods/domain"

	"gorm.io/gorm"
)

func MigrateGoods(db *gorm.DB) error {
	// Migrate the schema
	return db.AutoMigrate(&domain.Goods{})
}

func SeedGoods(db *gorm.DB) error {
	// Create a new goods
	newGoods := domain.Goods{
		Name:        "Test Product",
		CategoryID:  1,
		Price:       100.00,
		Stock:       10,
		Description: "This is a test product.",
		ImageURL:    "http://example.com/testproduct.jpg",
	}

	// Create the goods
	if err := db.Create(&newGoods).Error; err != nil {
		return err
	}

	return nil
}