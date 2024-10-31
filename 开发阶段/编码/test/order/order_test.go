package order

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"order/domain"
	"order/repository"
	"order/service"
	"gorm.io/gorm"
)

type MockOrderRepository struct {
	db *gorm.DB
}

func (m *MockOrderRepository) CreateOrder(ctx context.Context, order *domain.Order) error {
	return nil
}

func (m *MockOrderRepository) GetOrderByID(ctx context.Context, id int) (*domain.Order, error) {
	if id == 1 {
		return &domain.Order{
			ID:        1,
			UserID:    1,
			GoodsID:   1,
			Quantity:  2,
			Price:     200.00,
			Status:    1, // Pending
			CreatedAt: time.Now(),
		}, nil
	}
	return nil, gorm.ErrRecordNotFound
}

func (m *MockOrderRepository) ListOrders(ctx context.Context, userId int, page, limit int) ([]*domain.Order, error) {
	return []*domain.Order{
		{
			ID:        1,
			UserID:    1,
			GoodsID:   1,
			Quantity:  2,
			Price:     200.00,
			Status:    1, // Pending
			CreatedAt: time.Now(),
		},
	}, nil
}

func (m *MockOrderRepository) UpdateOrderStatus(ctx context.Context, id int, status int) error {
	return nil
}

func (m *MockOrderRepository) CancelOrder(ctx context.Context, id int) error {
	return nil
}

func TestOrderService_CreateOrder(t *testing.T) {
	db, mock, err := SetupTestDB()
	if err != nil {
		t.Fatalf("Failed to set up test DB: %v", err)
	}
	defer db.Close()

	mock.On("CreateOrder", mock.Anything, mock.Anything).Return(nil)

	orderService := service.NewOrderService(mock)
	ctx := context.Background()

	order := &domain.Order{
		UserID: 1,
		GoodsID: 1,
		Quantity: 2,
		Price: 200.00,
	}

	err = orderService.CreateOrder(ctx, order)
	assert.NoError(t, err)
	assert.NotNil(t, order)
	assert.Equal(t, 1, order.ID)
}