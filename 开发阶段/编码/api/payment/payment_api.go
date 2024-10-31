package api

import (
	"context"
	"encoding/json"
	"net/http"
	"payment/domain"
	"payment/service"

	"github.com/gin-gonic/gin"
)

type PaymentServiceInterface interface {
	ProcessPayment(ctx context.Context, orderId int, paymentDetails *domain.PaymentDetails) error
	GetPaymentStatus(ctx context.Context, orderId int) (*domain.PaymentStatus, error)
	RefundPayment(ctx context.Context, orderId int, refundAmount decimal.Decimal) error
	GetPaymentHistory(ctx context.Context, userId int) ([]*domain.PaymentHistory, error)
}

func (h *Handler) ProcessPayment(c *gin.Context) {
	ctx := c.Request.Context()
	var paymentDetails domain.PaymentDetails

	if err := c.ShouldBindJSON(&paymentDetails); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orderId := c.Param("orderId")

	if err := h.paymentService.ProcessPayment(ctx, orderId, &paymentDetails); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Payment processed successfully"})
}

func (h *Handler) GetPaymentStatus(c *gin.Context) {
	ctx := c.Request.Context()
	orderId := c.Param("orderId")

	paymentStatus, err := h.paymentService.GetPaymentStatus(ctx, orderId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, paymentStatus)
}

func (h *Handler) RefundPayment(c *gin.Context) {
	ctx := c.Request.Context()
	var refundRequest domain.RefundRequest

	if err := c.ShouldBindJSON(&refundRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orderId := c.Param("orderId")

	if err := h.paymentService.RefundPayment(ctx, orderId, refundRequest.Amount); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Refund processed successfully"})
}

func (h *Handler) GetPaymentHistory(c *gin.Context) {
	ctx := c.Request.Context()
	userId := c.Param("userId")

	paymentHistory, err := h.paymentService.GetPaymentHistory(ctx, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, paymentHistory)
}