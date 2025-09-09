//go:build slow
// +build slow

package server_test

import (
	"context"
	"testing"
	"time"
)

func TestMockServer_GracefulShutdown(t *testing.T) {
	s := NewMockServer()

	go func() {
		if err := s.Start(); err != nil {
			t.Errorf("server error: %v", err)
		}
	}()

	time.Sleep(200 * time.Millisecond)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err := s.Shutdown(ctx)
	if err != nil {
		t.Errorf("shutdown failed: %v", err)
	}
}

func TestMockServer_RequestDuringShutdown(t *testing.T) {
	mock := NewMockServer()

	go func() {
		if err := mock.Start(); err != nil {
			t.Errorf("server error: %v", err)
		}
	}()

	time.Sleep(200 * time.Millisecond)

	if ms, ok := mock.(*mockServer); ok {
		ms.HandleRequest(3 * time.Second)
	}

	time.Sleep(1 * time.Second)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	start := time.Now()
	err := mock.Shutdown(ctx)
	duration := time.Since(start)

	if err != nil {
		t.Errorf("shutdown failed: %v", err)
	}

	if duration < 2*time.Second {
		t.Errorf("expected graceful shutdown to wait for requests, but took %v", duration)
	} else {
		t.Logf("graceful shutdown took %v as expected", duration)
	}
}
