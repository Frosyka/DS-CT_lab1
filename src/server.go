package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func readyHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
	log.Println("/ready has been called. Server is ok")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "READY")
	log.Println("/health has been called. Server is ready")
}

func gracefulShutdown(server *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v\n", err)
	}

	log.Println("Server exiting")
}

func main() {
	http.HandleFunc("/ready", readyHandler)
	http.HandleFunc("/health", healthHandler)

	server := &http.Server{Addr: ":8092"}

	go gracefulShutdown(server)

	log.Println("Starting server at port 8092")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("could not listen on %s: %v\n", server.Addr, err)
	}
}
