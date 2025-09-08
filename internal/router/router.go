package router

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(engine *gin.Engine) {

	for _, h := range getPlatformHandlers() {
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