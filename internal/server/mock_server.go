package server

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type mockServer struct {
	running bool
	mu      sync.Mutex
	done    chan struct{}

	activeRequests sync.WaitGroup
}

func NewMockServer() Server {
	return &mockServer{
		done: make(chan struct{}),
	}
}

func (m *mockServer) Start() error {
	m.mu.Lock()
	m.running = true
	m.mu.Unlock()

	fmt.Println("[mockServer] started")
	<-m.done
	fmt.Println("[mockServer] stopped")
	return nil
}

func (m *mockServer) Shutdown(ctx context.Context) error {
	m.mu.Lock()
	if !m.running {
		m.mu.Unlock()
		return nil
	}
	fmt.Println("[mockServer] shutdown requested...")
	m.mu.Unlock()

	doneCh := make(chan struct{})
	go func() {
		m.activeRequests.Wait()
		close(doneCh)
	}()

	select {
	case <-doneCh:
		fmt.Println("[mockServer] all requests finished")
	case <-ctx.Done():
		return ctx.Err()
	}

	select {
	case <-time.After(500 * time.Millisecond):
		close(m.done)
	case <-ctx.Done():
		return ctx.Err()
	}

	m.mu.Lock()
	m.running = false
	m.mu.Unlock()
	return nil
}

func (m *mockServer) HandleRequest(duration time.Duration) {
	m.activeRequests.Add(1)
	go func() {
		defer m.activeRequests.Done()
		fmt.Println("[mockServer] request started...")
		time.Sleep(duration)
		fmt.Println("[mockServer] request finished")
	}()
}
