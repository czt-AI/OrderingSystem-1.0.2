package service

import (
	"context"
	"goods/domain"
	"goods/repository"
)

type GoodsService struct {
	repo repository.GoodsRepository
}

func NewGoodsService(repo repository.GoodsRepository) *GoodsService {
	return &GoodsService{repo: repo}
}

func (s *GoodsService) GetGoodsByID(ctx context.Context, id int) (*domain.Goods, error) {
	return s.repo.GetGoodsByID(ctx, id)
}

func (s *GoodsService) ListGoods(ctx context.Context, categoryID int, page, limit int) ([]*domain.Goods, error) {
	return s.repo.ListGoods(ctx, categoryID, page, limit)
}

func (s *GoodsService) CreateGoods(ctx context.Context, goods *domain.Goods) error {
	return s.repo.CreateGoods(ctx, goods)
}

func (s *GoodsService) UpdateGoods(ctx context.Context, id int, goods *domain.Goods) error {
	return s.repo.UpdateGoods(ctx, id, goods)
}

func (s *GoodsService) DeleteGoods(ctx context.Context, id int) error {
	return s.repo.DeleteGoods(ctx, id)
}