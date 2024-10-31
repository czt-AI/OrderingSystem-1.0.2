package db

import (
	"database/sql"
	"fmt"
	"payment/domain"

	"gorm.io/gorm"
)

func MigratePayment(db *gorm.DB) error {
	// Migrate the schema
	return db.AutoMigrate(&domain.Payment{})
}

func SeedPayment(db *gorm.DB) error {
	// Create a new payment
	newPayment := domain.Payment{
		OrderID:    1,
		PaymentMethod: "credit_card",
		PaymentTime:  time.Now(),
		PaymentStatus: 1, // 1: Pending
		Amount:      200.00,
	}

	// Create the payment
	if err := db.Create(&newPayment).Error; err != nil {
		return err
	}

	return nil
}