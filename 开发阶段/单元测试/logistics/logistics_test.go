package logistics

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"logistics/domain"
)

func TestShipmentTrackingValidation(t *testing.T) {
	// Test valid shipment tracking
	shipmentTracking := domain.ShipmentTracking{
		TrackingNumber: "LN123456789",
		Status:         "in_transit",
		EstimatedDelivery: time.Now().AddDate(0, 0, 5), // 5 days from now
	}

	err := shipmentTracking.Validate()
	assert.NoError(t, err)

	// Test invalid tracking number
	shipmentTracking.TrackingNumber = ""
	err = shipmentTracking.Validate()
	assert.Error(t, err)

	// Test invalid status
	shipmentTracking.Status = ""
	err = shipmentTracking.Validate()
	assert.Error(t, err)

	// Test invalid estimated delivery
	shipmentTracking.EstimatedDelivery = time.Time{}
	err = shipmentTracking.Validate()
	assert.Error(t, err)
}