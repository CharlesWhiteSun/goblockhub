package handler

import (
	"goblockhub/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BinanceHandler struct {
	svc service.IPlatformService
}

func NewBinanceHandler(svc service.IPlatformService) IPlatformHandler {
	return &BinanceHandler{svc: svc}
}

func (h *BinanceHandler) RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api/binance")
	api.GET("/status", h.getStatus)
}

func (h *BinanceHandler) getStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": h.svc.GetStatus()})
}
