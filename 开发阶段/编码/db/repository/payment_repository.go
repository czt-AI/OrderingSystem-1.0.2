package repository

import (
	"payment/domain"

	"gorm.io/gorm"
)

type PaymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) *PaymentRepository {
	return &PaymentRepository{db: db}
}

func (r *PaymentRepository) ProcessPayment(ctx context.Context, orderId int, paymentDetails *domain.PaymentDetails) error {
	// Process the payment logic here
	// This is a placeholder for the actual payment processing code
	return nil
}

func (r *PaymentRepository) GetPaymentStatus(ctx context.Context, orderId int) (*domain.PaymentStatus, error) {
	var paymentStatus domain.PaymentStatus
	if err := r.db.WithContext(ctx).First(&paymentStatus, "order_id = ?", orderId).Error; err != nil {
		return nil, err
	}
	return &paymentStatus, nil
}

func (r *PaymentRepository) RefundPayment(ctx context.Context, orderId int, refundAmount decimal.Decimal) error {
	// Refund logic here
	// This is a placeholder for the actual refund code
	return nil
}

func (r *PaymentRepository) GetPaymentHistory(ctx context.Context, userId int) ([]*domain.PaymentHistory, error) {
	var paymentHistory []*domain.PaymentHistory
	if err := r.db.WithContext(ctx).Where("user_id = ?", userId).Find(&paymentHistory).Error; err != nil {
		return nil, err
	}
	return paymentHistory, nil
}