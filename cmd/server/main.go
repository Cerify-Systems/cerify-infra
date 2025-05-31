package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/cerify/cerify-infra/internal/app"
	"github.com/cerify/cerify-infra/internal/config"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found")
	}

	// Initialize configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Create application context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Initialize and start the application
	application, err := app.New(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}

	// Start the application
	go func() {
		if err := application.Start(); err != nil {
			log.Printf("Error starting application: %v", err)
			cancel()
		}
	}()

	// Handle graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")
	if err := application.Shutdown(ctx); err != nil {
		log.Printf("Error during shutdown: %v", err)
	}
} 