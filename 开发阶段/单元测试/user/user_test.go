package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"user/domain"
)

func TestUserValidation(t *testing.T) {
	// Test valid user
	user := domain.User{
		Username: "testuser",
		Password: "password123",
		Email:    "testuser@example.com",
		Phone:    "1234567890",
	}

	err := user.Validate()
	assert.NoError(t, err)

	// Test invalid username
	user.Username = ""
	err = user.Validate()
	assert.Error(t, err)

	// Test invalid password
	user.Password = ""
	err = user.Validate()
	assert.Error(t, err)

	// Test invalid email
	user.Email = "invalid-email"
	err = user.Validate()
	assert.Error(t, err)

	// Test invalid phone
	user.Phone = "invalid-phone"
	err = user.Validate()
	assert.Error(t, err)
}