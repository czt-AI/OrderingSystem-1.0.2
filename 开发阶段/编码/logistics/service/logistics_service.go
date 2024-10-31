package service

import (
	"context"
	"logistics/domain"
	"logistics/repository"
)

type LogisticsService struct {
	repo repository.LogisticsRepository
}

func NewLogisticsService(repo repository.LogisticsRepository) *LogisticsService {
	return &LogisticsService{repo: repo}
}

func (s *LogisticsService) TrackShipment(ctx context.Context, trackingNumber string) (*domain.ShipmentTracking, error) {
	return s.repo.TrackShipment(ctx, trackingNumber)
}

func (s *LogisticsService) GetShipmentDetails(ctx context.Context, trackingNumber string) (*domain.ShipmentDetails, error) {
	return s.repo.GetShipmentDetails(ctx, trackingNumber)
}

func (s *LogisticsService) UpdateShipmentStatus(ctx context.Context, trackingNumber string, status domain.ShipmentStatus) error {
	return s.repo.UpdateShipmentStatus(ctx, trackingNumber, status)
}

func (s *LogisticsService) GetShipmentHistory(ctx context.Context, userId int) ([]*domain.ShipmentHistory, error) {
	return s.repo.GetShipmentHistory(ctx, userId)
}