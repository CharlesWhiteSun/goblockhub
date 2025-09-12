package router

import (
	"goblockhub/internal/handler"
	"goblockhub/internal/response"
	"goblockhub/internal/service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(engine *gin.Engine) {

	respHandler := response.NewResponseHandler()

	handlers := []handler.IPlatformHandler{
		handler.NewBinanceHandler(service.NewBinanceService(), respHandler),
		handler.NewOKXHandler(service.NewOKXService(), respHandler),
	}

	for _, h := range handlers {
		h.RegisterRoutes(engine)
	}

	engine.GET("/slow", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "finished slow API")
	})

	engine.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
}
