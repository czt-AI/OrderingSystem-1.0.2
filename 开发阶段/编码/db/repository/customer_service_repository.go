package repository

import (
	"customer_service/domain"

	"gorm.io/gorm"
)

type CustomerServiceRepository struct {
	db *gorm.DB
}

func NewCustomerServiceRepository(db *gorm.DB) *CustomerServiceRepository {
	return &CustomerServiceRepository{db: db}
}

func (r *CustomerServiceRepository) SubmitCustomerFeedback(ctx context.Context, feedback *domain.CustomerFeedback) error {
	return r.db.WithContext(ctx).Create(feedback).Error
}

func (r *CustomerServiceRepository) GetCustomerFeedback(ctx context.Context, userId int) ([]*domain.CustomerFeedback, error) {
	var feedbacks []*domain.CustomerFeedback
	if err := r.db.WithContext(ctx).Where("user_id = ?", userId).Find(&feedbacks).Error; err != nil {
		return nil, err
	}
	return feedbacks, nil
}

func (r *CustomerServiceRepository) GetCustomerSupportTickets(ctx context.Context, userId int) ([]*domain.CustomerSupportTicket, error) {
	var tickets []*domain.CustomerSupportTicket
	if err := r.db.WithContext(ctx).Where("user_id = ?", userId).Find(&tickets).Error; err != nil {
		return nil, err
	}
	return tickets, nil
}

func (r *CustomerServiceRepository) ResolveCustomerSupportTicket(ctx context.Context, ticketId int) error {
	return r.db.WithContext(ctx).Delete(&domain.CustomerSupportTicket{}, "id = ?", ticketId).Error
}