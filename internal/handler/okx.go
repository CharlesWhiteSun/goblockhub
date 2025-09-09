package handler

import (
	"goblockhub/internal/response"
	"goblockhub/internal/service"

	"github.com/gin-gonic/gin"
)

type OKXHandler struct {
	svc service.IPlatformService
	resp response.IResponseHandler
}

func NewOKXHandler(svc service.IPlatformService, resp response.IResponseHandler) IPlatformHandler {
	return &OKXHandler{svc: svc, resp: resp}
}

func (h *OKXHandler) RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api/okx")
	api.GET("/status", h.getStatus)
}

func (h *OKXHandler) getStatus(c *gin.Context) {
	h.resp.Success(c, gin.H{"status": h.svc.GetStatus()}, "OKX status retrieved successfully")
}
