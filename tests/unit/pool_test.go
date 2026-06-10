package unit


import (
	"net/http/httputil"
	"testing"
	"hydraload/internal/backend"
)


func TestPoolAddBackend(t *testing.T){
	
	pool := backend.NewPool()

	b, err := backend.NewBackend("http://localhost:9001", 1, &httputil.ReverseProxy{})
	if err != nil {
		t.Fatalf("Failed to create backend: %v", err)
	}

	if err != nil {
		t.Fatalf("Failed to add backend: %v", err)
	}
	err = pool.AddBackend(b)

	if err != nil {
		t.Fatalf("Failed to add backend: %v", err)
	}
	if pool.Size() != 1 {
		t.Fatalf(
			"expected 1 got %d",
			pool.Size(),
		)
	}
}