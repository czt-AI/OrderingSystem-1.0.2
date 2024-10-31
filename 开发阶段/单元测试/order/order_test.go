package order

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"order/domain"
)

func TestOrderValidation(t *testing.T) {
	// Test valid order
	order := domain.Order{
		UserID:   1,
		GoodsID:  1,
		Quantity: 2,
		Price:    200.00,
		Status:   1, // Pending
	}

	err := order.Validate()
	assert.NoError(t, err)

	// Test invalid user ID
	order.UserID = 0
	err = order.Validate()
	assert.Error(t, err)

	// Test invalid goods ID
	order.GoodsID = 0
	err = order.Validate()
	assert.Error(t, err)

	// Test invalid quantity
	order.Quantity = 0
	err = order.Validate()
	assert.Error(t, err)

	// Test invalid price
	order.Price = -1.00
	err = order.Validate()
	assert.Error(t, err)
}