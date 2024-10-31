package api

import (
	"context"
	"encoding/json"
	"net/http"
	"goods/domain"
	"goods/service"
	"goods/validation"

	"github.com/gin-gonic/gin"
)

type GoodsServiceInterface interface {
	GetGoodsByID(ctx context.Context, id int) (*domain.Goods, error)
	ListGoods(ctx context.Context, categoryID int, page, limit int) ([]*domain.Goods, error)
	CreateGoods(ctx context.Context, goods *domain.Goods) error
	UpdateGoods(ctx context.Context, id int, goods *domain.Goods) error
	DeleteGoods(ctx context.Context, id int) error
}

func (h *Handler) GetGoodsByID(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")

	goods, err := h.goodsService.GetGoodsByID(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, goods)
}

func (h *Handler) ListGoods(c *gin.Context) {
	ctx := c.Request.Context()
	categoryID := c.Query("categoryID")
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")

	goodsList, err := h.goodsService.ListGoods(ctx, categoryID, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, goodsList)
}

func (h *Handler) CreateGoods(c *gin.Context) {
	ctx := c.Request.Context()
	var goods domain.Goods

	if err := c.ShouldBindJSON(&goods); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validation.ValidateGoods(&goods); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.goodsService.CreateGoods(ctx, &goods); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Goods created successfully"})
}

func (h *Handler) UpdateGoods(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")
	var goods domain.Goods

	if err := c.ShouldBindJSON(&goods); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.goodsService.UpdateGoods(ctx, id, &goods); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Goods updated successfully"})
}

func (h *Handler) DeleteGoods(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")

	if err := h.goodsService.DeleteGoods(ctx, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Goods deleted successfully"})
}