package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

type Server interface {
	Start() error
	Shutdown(ctx context.Context) error
}

type ginServer struct {
	httpServer *http.Server
}

type RouterSetupFn func(engine *gin.Engine)

func NewGinServer(addr string, setup RouterSetupFn) Server {
	engine := gin.Default()

	if setup != nil {
		setup(engine)
	}

	return &ginServer{
		httpServer: &http.Server{
			Addr:    addr,
			Handler: engine,
		},
	}
}

func (s *ginServer) Start() error {
	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Server error: %v\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("Shutdown signal received...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return s.Shutdown(ctx)
}

func (s *ginServer) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
