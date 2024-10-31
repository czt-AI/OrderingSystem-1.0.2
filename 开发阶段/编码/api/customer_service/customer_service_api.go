package api

import (
	"context"
	"encoding/json"
	"net/http"
	"customer_service/domain"
	"customer_service/service"

	"github.com/gin-gonic/gin"
)

type CustomerServiceServiceInterface interface {
	SubmitCustomerFeedback(ctx context.Context, feedback *domain.CustomerFeedback) error
	GetCustomerFeedback(ctx context.Context, userId int) ([]*domain.CustomerFeedback, error)
	GetCustomerSupportTickets(ctx context.Context, userId int) ([]*domain.CustomerSupportTicket, error)
	ResolveCustomerSupportTicket(ctx context.Context, ticketId int) error
}

func (h *Handler) SubmitCustomerFeedback(c *gin.Context) {
	ctx := c.Request.Context()
	var feedback domain.CustomerFeedback

	if err := c.ShouldBindJSON(&feedback); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.customerServiceService.SubmitCustomerFeedback(ctx, &feedback); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Customer feedback submitted successfully"})
}

func (h *Handler) GetCustomerFeedback(c *gin.Context) {
	ctx := c.Request.Context()
	userId := c.Param("userId")

	feedbacks, err := h.customerServiceService.GetCustomerFeedback(ctx, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, feedbacks)
}

func (h *Handler) GetCustomerSupportTickets(c *gin.Context) {
	ctx := c.Request.Context()
	userId := c.Param("userId")

	tickets, err := h.customerServiceService.GetCustomerSupportTickets(ctx, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tickets)
}

func (h *Handler) ResolveCustomerSupportTicket(c *gin.Context) {
	ctx := c.Request.Context()
	ticketId := c.Param("ticketId")

	if err := h.customerServiceService.ResolveCustomerSupportTicket(ctx, ticketId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Customer support ticket resolved successfully"})
}