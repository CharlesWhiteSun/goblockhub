package handler

import (
	"goblockhub/internal/response"
	"goblockhub/internal/service"

	"github.com/gin-gonic/gin"
)

type BinanceHandler struct {
	svc service.IPlatformService
	resp response.IResponseHandler
}

func NewBinanceHandler(svc service.IPlatformService, resp response.IResponseHandler) IPlatformHandler {
	return &BinanceHandler{svc: svc, resp: resp}
}

func (h *BinanceHandler) RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api/binance")
	api.GET("/status", h.getStatus)
}

func (h *BinanceHandler) getStatus(c *gin.Context) {
	h.resp.Success(c, gin.H{"status": h.svc.GetStatus()}, "Binance status retrieved successfully")
}
