package server

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server

}

func NewServer(addr string) *Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(
		w http.ResponseWriter,  
		r *http.Request,
	){
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}) 

	httpServer := &http.Server {
		Addr: addr,
		Handler: mux,
		ReadTimeout: 10 *time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout: 120 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		MaxHeaderBytes: 1 << 20, 
	}
	// readtimeout -> prevent slowloris attack
	// writetimeout -> prevent hanging connections
	// idletimeout -> close idle connections
	// readheadertimeout -> Protects from malicious clients.
	// maxheaderbytes -> limit size of request headers

	return &Server {
		httpServer: httpServer,
	}
}

func (s *Server ) Start() error {
	return s.httpServer.ListenAndServer()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}