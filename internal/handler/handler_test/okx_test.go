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

func TestOKXHandler_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	respHandler := response.NewResponseHandler()
	svc := service.NewOKXService()
	h := handler.NewOKXHandler(svc, respHandler)

	r := gin.New()
	h.RegisterRoutes(r)

	req := httptest.NewRequest("GET", "/api/okx/status", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), `"success":true`)
	assert.Contains(t, w.Body.String(), `"code":1`)
	assert.Contains(t, w.Body.String(), `"status":"OKX OK"`)
}

func TestOKXHandler_Error(t *testing.T) {
	gin.SetMode(gin.TestMode)
	respHandler := response.NewResponseHandler()

	r := gin.New()
	r.GET("/api/okx/error", func(c *gin.Context) {
		respHandler.Error(c, errorx.NOT_FOUND, "Not found test")
	})

	req := httptest.NewRequest("GET", "/api/okx/error", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), `"success":false`)
	assert.Contains(t, w.Body.String(), `"code":10003`)
	assert.Contains(t, w.Body.String(), `"message":"Not found test"`)
}
