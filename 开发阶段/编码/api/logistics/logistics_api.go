package api

import (
	"context"
	"encoding/json"
	"net/http"
	"logistics/domain"
	"logistics/service"

	"github.com/gin-gonic/gin"
)

type LogisticsServiceInterface interface {
	TrackShipment(ctx context.Context, trackingNumber string) (*domain.ShipmentTracking, error)
	GetShipmentDetails(ctx context.Context, trackingNumber string) (*domain.ShipmentDetails, error)
	UpdateShipmentStatus(ctx context.Context, trackingNumber string, status domain.ShipmentStatus) error
	GetShipmentHistory(ctx context.Context, userId int) ([]*domain.ShipmentHistory, error)
}

func (h *Handler) TrackShipment(c *gin.Context) {
	ctx := c.Request.Context()
	trackingNumber := c.Param("trackingNumber")

	shipmentTracking, err := h.logisticsService.TrackShipment(ctx, trackingNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, shipmentTracking)
}

func (h *Handler) GetShipmentDetails(c *gin.Context) {
	ctx := c.Request.Context()
	trackingNumber := c.Param("trackingNumber")

	shipmentDetails, err := h.logisticsService.GetShipmentDetails(ctx, trackingNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, shipmentDetails)
}

func (h *Handler) UpdateShipmentStatus(c *gin.Context) {
	ctx := c.Request.Context()
	trackingNumber := c.Param("trackingNumber")
	status := c.DefaultQuery("status", "in_transit")

	if err := h.logisticsService.UpdateShipmentStatus(ctx, trackingNumber, status); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Shipment status updated successfully"})
}

func (h *Handler) GetShipmentHistory(c *gin.Context) {
	ctx := c.Request.Context()
	userId := c.Param("userId")

	shipmentHistory, err := h.logisticsService.GetShipmentHistory(ctx, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, shipmentHistory)
}