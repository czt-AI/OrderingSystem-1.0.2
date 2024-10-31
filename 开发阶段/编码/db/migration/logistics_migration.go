package db

import (
	"database/sql"
	"fmt"
	"logistics/domain"

	"gorm.io/gorm"
)

func MigrateLogistics(db *gorm.DB) error {
	// Migrate the schema
	return db.AutoMigrate(&domain.Shipment{})
}

func SeedLogistics(db *gorm.DB) error {
	// Create a new shipment
	newShipment := domain.Shipment{
		OrderID:          1,
		TrackingNumber:   "LN123456789",
		LogisticsCompany: "Example Logistics",
		EstimatedDelivery: time.Now().AddDate(0, 0, 5), // 5 days from now
	}

	// Create the shipment
	if err := db.Create(&newShipment).Error; err != nil {
		return err
	}

	return nil
}