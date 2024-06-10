package main

import (
	"log"

	"github.com/guilherme-luvi/go-queue-worker/internal/config"
	"github.com/guilherme-luvi/go-queue-worker/internal/queue"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize Kafka consumer
	consumer, err := queue.NewConsumer(cfg)
	if err != nil {
		log.Fatalf("Failed to create Kafka consumer: %v", err)
	}

	// Start consuming messages
	err = consumer.Consume(cfg)
	if err != nil {
		log.Fatalf("Error while consuming messages: %v", err)
	}
}
