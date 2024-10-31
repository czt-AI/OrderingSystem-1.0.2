package goods

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"goods/domain"
)

func TestGoodsValidation(t *testing.T) {
	// Test valid goods
	goods := domain.Goods{
		Name:        "Test Product",
		CategoryID:  1,
		Price:       100.00,
		Stock:       10,
		Description: "This is a test product.",
		ImageURL:    "http://example.com/testproduct.jpg",
	}

	err := goods.Validate()
	assert.NoError(t, err)

	// Test invalid name
	goods.Name = ""
	err = goods.Validate()
	assert.Error(t, err)

	// Test invalid price
	goods.Price = -1.00
	err = goods.Validate()
	assert.Error(t, err)

	// Test invalid stock
	goods.Stock = -1
	err = goods.Validate()
	assert.Error(t, err)
}