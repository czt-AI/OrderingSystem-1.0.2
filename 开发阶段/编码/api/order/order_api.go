package api

import (
	"context"
	"encoding/json"
	"net/http"
	"order/domain"
	"order/service"

	"github.com/gin-gonic/gin"
)

type OrderServiceInterface interface {
	CreateOrder(ctx context.Context, order *domain.Order) error
	GetOrderByID(ctx context.Context, id int) (*domain.Order, error)
	ListOrders(ctx context.Context, userId int, page, limit int) ([]*domain.Order, error)
	UpdateOrderStatus(ctx context.Context, id int, status int) error
	CancelOrder(ctx context.Context, id int) error
}

func (h *Handler) CreateOrder(c *gin.Context) {
	ctx := c.Request.Context()
	var order domain.Order

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.orderService.CreateOrder(ctx, &order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Order created successfully"})
}

func (h *Handler) GetOrderByID(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")

	order, err := h.orderService.GetOrderByID(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, order)
}

func (h *Handler) ListOrders(c *gin.Context) {
	ctx := c.Request.Context()
	userId := c.Query("userId")
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")

	orders, err := h.orderService.ListOrders(ctx, userId, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orders)
}

func (h *Handler) UpdateOrderStatus(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")
	status := c.DefaultQuery("status", "1")

	if err := h.orderService.UpdateOrderStatus(ctx, id, status); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order status updated successfully"})
}

func (h *Handler) CancelOrder(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")

	if err := h.orderService.CancelOrder(ctx, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order canceled successfully"})
}