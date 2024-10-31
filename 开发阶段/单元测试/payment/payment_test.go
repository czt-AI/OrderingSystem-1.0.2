package payment

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"payment/domain"
)

func TestPaymentValidation(t *testing.T) {
	// Test valid payment details
	paymentDetails := domain.PaymentDetails{
		OrderID:    1,
		PaymentMethod: "credit_card",
		Amount:      200.00,
	}

	err := paymentDetails.Validate()
	assert.NoError(t, err)

	// Test invalid order ID
	paymentDetails.OrderID = 0
	err = paymentDetails.Validate()
	assert.Error(t, err)

	// Test invalid payment method
	paymentDetails.PaymentMethod = ""
	err = paymentDetails.Validate()
	assert.Error(t, err)

	// Test invalid amount
	paymentDetails.Amount = -1.00
	err = paymentDetails.Validate()
	assert.Error(t, err)
}