package goods

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"goods/domain"
	"goods/repository"
	"goods/service"
	"gorm.io/gorm"
)

type MockGoodsRepository struct {
	db *gorm.DB
}

func (m *MockGoodsRepository) GetGoodsByID(ctx context.Context, id int) (*domain.Goods, error) {
	if id == 1 {
		return &domain.Goods{
			ID:          1,
			Name:        "Test Product",
			CategoryID:  1,
			Price:       100.00,
			Stock:       10,
			Description: "This is a test product.",
			ImageURL:    "http://example.com/testproduct.jpg",
		}, nil
	}
	return nil, gorm.ErrRecordNotFound
}

func (m *MockGoodsRepository) ListGoods(ctx context.Context, categoryID int, page, limit int) ([]*domain.Goods, error) {
	return []*domain.Goods{
		{
			ID:          1,
			Name:        "Test Product",
			CategoryID:  1,
			Price:       100.00,
			Stock:       10,
			Description: "This is a test product.",
			ImageURL:    "http://example.com/testproduct.jpg",
		},
	}, nil
}

func (m *MockGoodsRepository) CreateGoods(ctx context.Context, goods *domain.Goods) error {
	return nil
}

func (m *MockGoodsRepository) UpdateGoods(ctx context.Context, id int, goods *domain.Goods) error {
	return nil
}

func (m *MockGoodsRepository) DeleteGoods(ctx context.Context, id int) error {
	return nil
}

func TestGoodsService_GetGoodsByID(t *testing.T) {
	db, mock, err := SetupTestDB()
	if err != nil {
		t.Fatalf("Failed to set up test DB: %v", err)
	}
	defer db.Close()

	mock.On("GetGoodsByID", mock.Anything, mock.Anything).Return(&domain.Goods{ID: 1, Name: "Test Product"}, nil)

	goodsService := service.NewGoodsService(mock)
	ctx := context.Background()

	goods, err := goodsService.GetGoodsByID(ctx, 1)
	assert.NoError(t, err)
	assert.NotNil(t, goods)
	assert.Equal(t, "Test Product", goods.Name)
}