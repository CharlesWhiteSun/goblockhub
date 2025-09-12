package handler

import (
	"goblockhub/internal/response"
	"goblockhub/internal/service"

	"github.com/CharlesWhiteSun/gomodx/errorx"

	"github.com/gin-gonic/gin"
)

type OKXHandler struct {
	svcStat service.IStatusService
	resp    response.IResponseHandler
}

func NewOKXHandler(svc *service.OKXService, resp response.IResponseHandler) IPlatformHandler {
	return &OKXHandler{
		svcStat: svc,
		resp:    resp,
	}
}

func (o *OKXHandler) RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api/okx")

	v1 := api.Group("/v1")
	{
		v1.GET("/status", o.GetStatus)
	}
}

func (o *OKXHandler) GetStatus(c *gin.Context) {
	ok, err := o.svcStat.GetStatus()
	if !ok {
		o.resp.Error(c, errorx.API_REQ_FAILED, err.Error())
		return
	}
	if err != nil {
		o.resp.Success(c, nil, err.Error())
		return
	}
	o.resp.Success(c, nil, "OK")
}
