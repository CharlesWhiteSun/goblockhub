package handler

import (
	"goblockhub/internal/response"
	"goblockhub/internal/service"

	"github.com/CharlesWhiteSun/gomodx/errorx"
	"github.com/gin-gonic/gin"
)

type BinanceHandler struct {
	svc service.IPlatformService
	resp response.IResponseHandler
}

func NewBinanceHandler(svc service.IPlatformService, resp response.IResponseHandler) IPlatformHandler {
	return &BinanceHandler{svc: svc, resp: resp}
}

func (b *BinanceHandler) RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api/binance")
	api.GET("/status", b.GetStatus)
}

func (b *BinanceHandler) GetStatus(c *gin.Context) {
	ok, err := b.svc.GetStatus()
	if !ok {
		b.resp.Error(c, errorx.API_REQ_FAILED, err.Error())
		return
	}
	if err != nil {
		b.resp.Success(c, nil, err.Error())
		return
	}
	b.resp.Success(c, nil, "OK")
}
