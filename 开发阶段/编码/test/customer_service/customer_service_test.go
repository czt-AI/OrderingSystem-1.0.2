package customer_service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"customer_service/domain"
	"customer_service/repository"
	"customer_service/service"
	"gorm.io/gorm"
)

type MockCustomerServiceRepository struct {
	db *gorm.DB
}

func (m *MockCustomerServiceRepository) SubmitCustomerFeedback(ctx context.Context, feedback *domain.CustomerFeedback) error {
	return nil
}

func (m *MockCustomerServiceRepository) GetCustomerFeedback(ctx context.Context, userId int) ([]*domain.CustomerFeedback, error) {
	return []*domain.CustomerFeedback{
		{
			UserId:    1,
			Message:   "Great product!",
			FeedbackType: "positive",
		},
	}, nil
}

func (m *MockCustomerServiceRepository) GetCustomerSupportTickets(ctx context.Context, userId int) ([]*domain.CustomerSupportTicket, error) {
	return []*domain.CustomerSupportTicket{
		{
			UserId:    1,
			Subject:   "Product issue",
			Message:   "The product arrived damaged.",
			Status:    "open",
		},
	}, nil
}

func (m *MockCustomerServiceRepository) ResolveCustomerSupportTicket(ctx context.Context, ticketId int) error {
	return nil
}

func TestCustomerService_SubmitCustomerFeedback(t *testing.T) {
	db, mock, err := SetupTestDB()
	if err != nil {
		t.Fatalf("Failed to set up test DB: %v", err)
	}
	defer db.Close()

	mock.On("SubmitCustomerFeedback", mock.Anything, mock.Anything).Return(nil)

	customerService := service.NewCustomerService(mock)
	ctx := context.Background()

	feedback := &domain.CustomerFeedback{
		UserId:    1,
		Message:   "Great product!",
		FeedbackType: "positive",
	}

	err = customerService.SubmitCustomerFeedback(ctx, feedback)
	assert.NoError(t, err)
}

func TestCustomerService_GetCustomerFeedback(t *testing.T) {
	db, mock, err := SetupTestDB()
	if err != nil {
		t.Fatalf("Failed to set up test DB: %v", err)
	}
	defer db.Close()

	mock.On("GetCustomerFeedback", mock.Anything, mock.Anything).Return([]*domain.CustomerFeedback{
		{
			UserId:    1,
			Message:   "Great product!",
			FeedbackType: "positive",
		},
	}, nil)

	customerService := service.NewCustomerService(mock)
	ctx := context.Background()

	feedbacks, err := customerService.GetCustomerFeedback(ctx, 1)
	assert.NoError(t, err)
	assert.Len(t, feedbacks, 1)
}