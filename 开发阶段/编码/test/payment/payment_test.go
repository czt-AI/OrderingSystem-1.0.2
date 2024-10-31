package payment

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"payment/domain"
	"payment/repository"
	"payment/service"
	"gorm.io/gorm"
)

type MockPaymentRepository struct {
	db *gorm.DB
}

func (m *MockPaymentRepository) ProcessPayment(ctx context.Context, orderId int, paymentDetails *domain.PaymentDetails) error {
	return nil
}

func (m *MockPaymentRepository) GetPaymentStatus(ctx context.Context, orderId int) (*domain.PaymentStatus, error) {
	if orderId == 1 {
		return &domain.PaymentStatus{
			OrderID:    orderId,
			PaymentMethod: "credit_card",
			PaymentTime:  time.Now(),
			PaymentStatus: 1, // Pending
			Amount:      200.00,
		}, nil
	}
	return nil, gorm.ErrRecordNotFound
}

func (m *MockPaymentRepository) RefundPayment(ctx context.Context, orderId int, refundAmount decimal.Decimal) error {
	return nil
}

func (m *MockPaymentRepository) GetPaymentHistory(ctx context.Context, userId int) ([]*domain.PaymentHistory, error) {
	return []*domain.PaymentHistory{
		{
			OrderID:    1,
			PaymentMethod: "credit_card",
			PaymentTime:  time.Now(),
			PaymentStatus: 1, // Pending
			Amount:      200.00,
		},
	}, nil
}

func TestPaymentService_ProcessPayment(t *testing.T) {
	db, mock, err := SetupTestDB()
	if err != nil {
		t.Fatalf("Failed to set up test DB: %v", err)
	}
	defer db.Close()

	mock.On("ProcessPayment", mock.Anything, mock.Anything).Return(nil)

	paymentService := service.NewPaymentService(mock)
	ctx := context.Background()

	paymentDetails := &domain.PaymentDetails{
		OrderID:    1,
		PaymentMethod: "credit_card",
		Amount:      200.00,
	}

	err = paymentService.ProcessPayment(ctx, 1, paymentDetails)
	assert.NoError(t, err)
}