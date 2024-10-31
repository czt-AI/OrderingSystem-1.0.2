package logistics

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"logistics/domain"
	"logistics/repository"
	"logistics/service"
	"gorm.io/gorm"
)

type MockLogisticsRepository struct {
	db *gorm.DB
}

func (m *MockLogisticsRepository) TrackShipment(ctx context.Context, trackingNumber string) (*domain.ShipmentTracking, error) {
	if trackingNumber == "LN123456789" {
		return &domain.ShipmentTracking{
			TrackingNumber: trackingNumber,
			Status:         "in_transit",
			EstimatedDelivery: time.Now().AddDate(0, 0, 5), // 5 days from now
		}, nil
	}
	return nil, gorm.ErrRecordNotFound
}

func (m *MockLogisticsRepository) GetShipmentDetails(ctx context.Context, trackingNumber string) (*domain.ShipmentDetails, error) {
	if trackingNumber == "LN123456789" {
		return &domain.ShipmentDetails{
			TrackingNumber: trackingNumber,
			Status:         "in_transit",
			LogisticsCompany: "Example Logistics",
		}, nil
	}
	return nil, gorm.ErrRecordNotFound
}

func (m *MockLogisticsRepository) UpdateShipmentStatus(ctx context.Context, trackingNumber string, status domain.ShipmentStatus) error {
	return nil
}

func (m *MockLogisticsRepository) GetShipmentHistory(ctx context.Context, userId int) ([]*domain.ShipmentHistory, error) {
	return []*domain.ShipmentHistory{
		{
			TrackingNumber: "LN123456789",
			Status:         "in_transit",
			LogisticsCompany: "Example Logistics",
		},
	}, nil
}

func TestLogisticsService_TrackShipment(t *testing.T) {
	db, mock, err := SetupTestDB()
	if err != nil {
		t.Fatalf("Failed to set up test DB: %v", err)
	}
	defer db.Close()

	mock.On("TrackShipment", mock.Anything, mock.Anything).Return(&domain.ShipmentTracking{
		TrackingNumber: "LN123456789",
		Status:         "in_transit",
		EstimatedDelivery: time.Now().AddDate(0, 0, 5), // 5 days from now
	}, nil)

	logisticsService := service.NewLogisticsService(mock)
	ctx := context.Background()

	shipmentTracking, err := logisticsService.TrackShipment(ctx, "LN123456789")
	assert.NoError(t, err)
	assert.NotNil(t, shipmentTracking)
	assert.Equal(t, "in_transit", shipmentTracking.Status)
}