package service

import (
	"context"
	"order/domain"
	"order/repository"
)

type OrderService struct {
	repo repository.OrderRepository
}

func NewOrderService(repo repository.OrderRepository) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) CreateOrder(ctx context.Context, order *domain.Order) error {
	return s.repo.CreateOrder(ctx, order)
}

func (s *OrderService) GetOrderByID(ctx context.Context, id int) (*domain.Order, error) {
	return s.repo.GetOrderByID(ctx, id)
}

func (s *OrderService) ListOrders(ctx context.Context, userId int, page, limit int) ([]*domain.Order, error) {
	return s.repo.ListOrders(ctx, userId, page, limit)
}

func (s *OrderService) UpdateOrderStatus(ctx context.Context, id int, status int) error {
	return s.repo.UpdateOrderStatus(ctx, id, status)
}

func (s *OrderService) CancelOrder(ctx context.Context, id int) error {
	return s.repo.CancelOrder(ctx, id)
}