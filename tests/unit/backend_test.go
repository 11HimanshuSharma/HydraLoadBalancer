package unit

import (
	"net/http/httputil"
	"testing"
	
	"hydraload/internal/backend"
)

func TestConnectionCounter(t *testing.T) {

	proxy := &httputil.ReverseProxy{}

	b, err := backend.NewBackend("http://localhost:9001", 1, proxy)
	if err != nil {
		t.Fatalf("Failed to create backend: %v", err)
	}
	b.IncrementConnections()
	b.IncrementConnections()

	if b.ActiveConnections() != 2 {
		t.Errorf("Expected 2 active connections, got %d", b.ActiveConnections())
	}
	b.DecrementConnections()
	if b.ActiveConnections() != 1 {
		t.Errorf("Expected 1 active connection, got %d", b.ActiveConnections())
	}	

	
}