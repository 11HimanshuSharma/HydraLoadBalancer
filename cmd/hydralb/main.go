package main


import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"hydraload/internal/app"
)

func main() {
	srv := server.NewServer(":8080")

	go func() {
		log.Println("HydraLB is starting on :8080")

		if err := srv.Start(); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	waitForShutdown(srv)
}

func waitForShutdown(srv *server.Server) {
	stop := make(chan os.Signal, 1)

	signal.Notify(
		stop,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	<-stop

	log.Println("Shutting down HydraLB...")

	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Failed to shutdown server: %v", err)
	}
	log.Println("HydraLB has been stopped.")
}