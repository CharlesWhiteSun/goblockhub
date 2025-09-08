package handler

import (
	"goblockhub/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OKXHandler struct {
	svc service.IPlatformService
}

func NewOKXHandler(svc service.IPlatformService) IPlatformHandler {
	return &OKXHandler{svc: svc}
}

func (h *OKXHandler) RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api/okx")
	api.GET("/status", h.getStatus)
}

func (h *OKXHandler) getStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": h.svc.GetStatus()})
}
