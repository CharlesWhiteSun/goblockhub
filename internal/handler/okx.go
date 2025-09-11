package handler

import (
	"goblockhub/internal/response"
	"goblockhub/internal/service"

	"github.com/CharlesWhiteSun/gomodx/errorx"

	"github.com/gin-gonic/gin"
)

type OKXHandler struct {
	svc service.IPlatformService
	resp response.IResponseHandler
}

func NewOKXHandler(svc service.IPlatformService, resp response.IResponseHandler) IPlatformHandler {
	return &OKXHandler{svc: svc, resp: resp}
}

func (o *OKXHandler) RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api/okx")
	api.GET("/status", o.getStatus)
}

func (o *OKXHandler) getStatus(c *gin.Context) {
	ok, err := o.svc.GetStatus()
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
