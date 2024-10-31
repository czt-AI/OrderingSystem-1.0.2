package repository

import (
	"order/domain"

	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) CreateOrder(ctx context.Context, order *domain.Order) error {
	return r.db.WithContext(ctx).Create(order).Error
}

func (r *OrderRepository) GetOrderByID(ctx context.Context, id int) (*domain.Order, error) {
	var order domain.Order
	if err := r.db.WithContext(ctx).First(&order, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *OrderRepository) ListOrders(ctx context.Context, userId int, page, limit int) ([]*domain.Order, error) {
	var orders []*domain.Order
	offset := (page - 1) * limit
	if err := r.db.WithContext(ctx).Where("user_id = ?", userId).Limit(limit).Offset(offset).Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *OrderRepository) UpdateOrderStatus(ctx context.Context, id int, status int) error {
	return r.db.WithContext(ctx).Model(&domain.Order{}).Where("id = ?", id).Updates(map[string]interface{}{"status": status}).Error
}

func (r *OrderRepository) CancelOrder(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Delete(&domain.Order{}, "id = ?", id).Error
}