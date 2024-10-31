package repository

import (
	"goods/domain"

	"gorm.io/gorm"
)

type GoodsRepository struct {
	db *gorm.DB
}

func NewGoodsRepository(db *gorm.DB) *GoodsRepository {
	return &GoodsRepository{db: db}
}

func (r *GoodsRepository) GetGoodsByID(ctx context.Context, id int) (*domain.Goods, error) {
	var goods domain.Goods
	if err := r.db.WithContext(ctx).First(&goods, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &goods, nil
}

func (r *GoodsRepository) ListGoods(ctx context.Context, categoryID int, page, limit int) ([]*domain.Goods, error) {
	var goods []*domain.Goods
	offset := (page - 1) * limit
	if err := r.db.WithContext(ctx).Where("category_id = ?", categoryID).Limit(limit).Offset(offset).Find(&goods).Error; err != nil {
		return nil, err
	}
	return goods, nil
}

func (r *GoodsRepository) CreateGoods(ctx context.Context, goods *domain.Goods) error {
	return r.db.WithContext(ctx).Create(goods).Error
}

func (r *GoodsRepository) UpdateGoods(ctx context.Context, id int, goods *domain.Goods) error {
	return r.db.WithContext(ctx).Model(&domain.Goods{}).Where("id = ?", id).Updates(goods).Error
}

func (r *GoodsRepository) DeleteGoods(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Delete(&domain.Goods{}, "id = ?", id).Error
}