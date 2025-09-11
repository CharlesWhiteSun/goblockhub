package handler_test

import (
	"goblockhub/internal/handler"
	"goblockhub/internal/response"
	"goblockhub/internal/service"
	"testing"

	"github.com/CharlesWhiteSun/gomodx/errorx"

	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestBinanceHandler_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	respHandler := response.NewResponseHandler()
	svc := service.NewBinanceService()
	h := handler.NewBinanceHandler(svc, respHandler)

	r := gin.New()
	h.RegisterRoutes(r)

	req := httptest.NewRequest("GET", "/api/binance/status", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), `"success":true`)
	assert.Contains(t, w.Body.String(), `"code":1`)
}

func TestBinanceHandler_Error(t *testing.T) {
	gin.SetMode(gin.TestMode)
	respHandler := response.NewResponseHandler()

	r := gin.New()
	r.GET("/api/binance/error", func(c *gin.Context) {
		respHandler.Error(c, errorx.INVALID_PARAMS, "Invalid params test")
	})

	req := httptest.NewRequest("GET", "/api/binance/error", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), `"success":false`)
	assert.Contains(t, w.Body.String(), `"code":10002`)
	assert.Contains(t, w.Body.String(), `"message":"Invalid params test"`)
}
