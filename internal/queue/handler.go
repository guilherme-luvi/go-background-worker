package queue

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/guilherme-luvi/go-queue-worker/internal/config"
	"github.com/guilherme-luvi/go-queue-worker/internal/database"
)

func handleMessage(cfg *config.Config, msg *kafka.Message) {
	// Process the message and update the cache
	err := database.UpdateCache(cfg, msg.Value)
	if err != nil {
		log.Printf("Failed to update cache: %v", err)
	}
}
