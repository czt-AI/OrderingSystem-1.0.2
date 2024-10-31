package service

import (
	"context"
	"payment/domain"
	"payment/repository"
)

type PaymentService struct {
	repo repository.PaymentRepository
}

func NewPaymentService(repo repository.PaymentRepository) *PaymentService {
	return &PaymentService{repo: repo}
}

func (s *PaymentService) ProcessPayment(ctx context.Context, orderId int, paymentDetails *domain.PaymentDetails) error {
	return s.repo.ProcessPayment(ctx, orderId, paymentDetails)
}

func (s *PaymentService) GetPaymentStatus(ctx context.Context, orderId int) (*domain.PaymentStatus, error) {
	return s.repo.GetPaymentStatus(ctx, orderId)
}

func (s *PaymentService) RefundPayment(ctx context.Context, orderId int, refundAmount decimal.Decimal) error {
	return s.repo.RefundPayment(ctx, orderId, refundAmount)
}

func (s *PaymentService) GetPaymentHistory(ctx context.Context, userId int) ([]*domain.PaymentHistory, error) {
	return s.repo.GetPaymentHistory(ctx, userId)
}