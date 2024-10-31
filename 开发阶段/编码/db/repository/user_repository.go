package repository

import (
	"user/domain"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetUserByID(ctx context.Context, id int) (*domain.User, error) {
	var user domain.User
	if err := r.db.WithContext(ctx).First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) CreateUser(ctx context.Context, user *domain.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *UserRepository) UpdateUser(ctx context.Context, id int, user *domain.User) error {
	return r.db.WithContext(ctx).Model(&domain.User{}).Where("id = ?", id).Updates(user).Error
}

func (r *UserRepository) DeleteUser(ctx context.Context, id int) error {
	return r.db.WithContext(ctx).Delete(&domain.User{}, "id = ?", id).Error
}