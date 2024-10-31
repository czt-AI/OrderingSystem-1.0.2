package customer_service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"customer_service/domain"
)

func TestCustomerFeedbackValidation(t *testing.T) {
	// Test valid customer feedback
	feedback := domain.CustomerFeedback{
		UserId:    1,
		Message:   "Great product!",
		FeedbackType: "positive",
	}

	err := feedback.Validate()
	assert.NoError(t, err)

	// Test invalid user ID
	feedback.UserId = 0
	err = feedback.Validate()
	assert.Error(t, err)

	// Test invalid message
	feedback.Message = ""
	err = feedback.Validate()
	assert.Error(t, err)

	// Test invalid feedback type
	feedback.FeedbackType = ""
	err = feedback.Validate()
	assert.Error(t, err)
}