package user

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"user/domain"
	"user/repository"
	"user/service"
	"gorm.io/gorm"
)

type MockUserRepository struct {
	db *gorm.DB
}

func (m *MockUserRepository) GetUserByID(ctx context.Context, id int) (*domain.User, error) {
	if id == 1 {
		return &domain.User{
			ID:       1,
			Username: "testuser",
			Password: "password123",
			Email:    "testuser@example.com",
			Phone:    "1234567890",
		}, nil
	}
	return nil, gorm.ErrRecordNotFound
}

func (m *MockUserRepository) CreateUser(ctx context.Context, user *domain.User) error {
	return nil
}

func (m *MockUserRepository) UpdateUser(ctx context.Context, id int, user *domain.User) error {
	return nil
}

func (m *MockUserRepository) DeleteUser(ctx context.Context, id int) error {
	return nil
}

func TestUserService_GetUserByID(t *testing.T) {
	db, mock, err := SetupTestDB()
	if err != nil {
		t.Fatalf("Failed to set up test DB: %v", err)
	}
	defer db.Close()

	mock.On(" GetUserByID", mock.Anything, mock.Anything).Return(&domain.User{ID: 1, Username: "testuser"}, nil)

	userService := service.NewUserService(mock)
	ctx := context.Background()

	user, err := userService.GetUserByID(ctx, 1)
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "testuser", user.Username)
}