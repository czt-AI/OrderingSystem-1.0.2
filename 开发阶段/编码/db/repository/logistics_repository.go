package repository

import (
	"logistics/domain"

	"gorm.io/gorm"
)

type LogisticsRepository struct {
	db *gorm.DB
}

func NewLogisticsRepository(db *gorm.DB) *LogisticsRepository {
	return &LogisticsRepository{db: db}
}

func (r *LogisticsRepository) TrackShipment(ctx context.Context, trackingNumber string) (*domain.ShipmentTracking, error) {
	var shipmentTracking domain.ShipmentTracking
	if err := r.db.WithContext(ctx).First(&shipmentTracking, "tracking_number = ?", trackingNumber).Error; err != nil {
		return nil, err
	}
	return &shipmentTracking, nil
}

func (r *LogisticsRepository) GetShipmentDetails(ctx context.Context, trackingNumber string) (*domain.ShipmentDetails, error) {
	var shipmentDetails domain.ShipmentDetails
	if err := r.db.WithContext(ctx).First(&shipmentDetails, "tracking_number = ?", trackingNumber).Error; err != nil {
		return nil, err
	}
	return &shipmentDetails, nil
}

func (r *LogisticsRepository) UpdateShipmentStatus(ctx context.Context, trackingNumber string, status domain.ShipmentStatus) error {
	return r.db.WithContext(ctx).Model(&domain.ShipmentDetails{}).Where("tracking_number = ?", trackingNumber).Updates(map[string]interface{}{"status": status}).Error
}

func (r *LogisticsRepository) GetShipmentHistory(ctx context.Context, userId int) ([]*domain.ShipmentHistory, error) {
	var shipmentHistory []*domain.ShipmentHistory
	if err := r.db.WithContext(ctx).Where("user_id = ?", userId).Find(&shipmentHistory).Error; err != nil {
		return nil, err
	}
	return shipmentHistory, nil
}