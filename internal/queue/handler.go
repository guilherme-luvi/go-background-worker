package queue

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/guilherme-luvi/go-queue-worker/internal/database"
	"github.com/guilherme-luvi/go-queue-worker/pkg/config"
)

func handleMessage(cfg *config.Config, msg *kafka.Message) error {
	// TODO: Implement the logic to handle the message

	// Update the cache with the message value
	err := database.UpdateCache(cfg, msg.Value)
	if err != nil {
		log.Printf("Failed to update cache: %v", err)
		return err
	}

	return nil
}
