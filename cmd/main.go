package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/guilherme-luvi/go-queue-worker/internal/queue"
	"github.com/guilherme-luvi/go-queue-worker/pkg/config"
)

func main() {
	// Load environment variables
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize Kafka consumer
	consumer, err := queue.NewConsumer(cfg)
	if err != nil {
		log.Fatalf("Failed to create Kafka consumer: %v", err)
	}

	// context to control the consumer goroutine and wait for the interrupt signal
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// Start consuming messages in a separate goroutine
	go func() {
		if err := consumer.Consume(ctx, cfg); err != nil {
			log.Fatalf("Error while consuming messages: %v", err)
		}
	}()

	// Wait for the interrupt signal to shutdown
	<-ctx.Done()
	log.Println("Shutting down...")

	consumer.Close()
}
