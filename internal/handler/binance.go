package handler

import (
	"goblockhub/internal/response"
	"goblockhub/internal/service"

	"github.com/CharlesWhiteSun/gomodx/errorx"
	"github.com/gin-gonic/gin"
)

type BinanceHandler struct {
	svcStat service.IStatusService
	svcTime service.ITimeService
	resp    response.IResponseHandler
}

func NewBinanceHandler(svc *service.BinanceService, resp response.IResponseHandler) IPlatformHandler {
	return &BinanceHandler{
		svcStat: svc,
		svcTime: svc,
		resp:    resp,
	}
}

func (b *BinanceHandler) RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api/binance")

	v1 := api.Group("/v1")
	{
		v1.GET("/status", b.getStat)
		v1.GET("/time", b.getTime)
	}
}

func (b *BinanceHandler) getStat(c *gin.Context) {
	ok, err := b.svcStat.GetStatus()
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

func (b *BinanceHandler) getTime(c *gin.Context) {
	ok, serverTime, err := b.svcTime.GetTime()
	if !ok {
		b.resp.Error(c, errorx.API_REQ_FAILED, err.Error())
		return
	}
	b.resp.Success(c, gin.H{"serverTime": serverTime}, "OK")
}
