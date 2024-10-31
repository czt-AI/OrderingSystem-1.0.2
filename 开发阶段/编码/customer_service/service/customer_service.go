package service

import (
	"context"
	"customer_service/domain"
	"customer_service/repository"
)

type CustomerService struct {
	repo repository.CustomerRepository
}

func NewCustomerService(repo repository.CustomerRepository) *CustomerService {
	return &CustomerService{repo: repo}
}

func (s *CustomerService) SubmitCustomerFeedback(ctx context.Context, feedback *domain.CustomerFeedback) error {
	return s.repo.SubmitCustomerFeedback(ctx, feedback)
}

func (s *CustomerService) GetCustomerFeedback(ctx context.Context, userId int) ([]*domain.CustomerFeedback, error) {
	return s.repo.GetCustomerFeedback(ctx, userId)
}

func (s *CustomerService) GetCustomerSupportTickets(ctx context.Context, userId int) ([]*domain.CustomerSupportTicket, error) {
	return s.repo.GetCustomerSupportTickets(ctx, userId)
}

func (s *CustomerService) ResolveCustomerSupportTicket(ctx context.Context, ticketId int) error {
	return s.repo.ResolveCustomerSupportTicket(ctx, ticketId)
}