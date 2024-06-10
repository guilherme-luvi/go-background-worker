package queue

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/guilherme-luvi/go-queue-worker/internal/config"
)

type Consumer struct {
	kafkaConsumer *kafka.Consumer
}

func NewConsumer(cfg *config.Config) (*Consumer, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": cfg.Kafka.Brokers,
		"group.id":          cfg.Kafka.GroupID,
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		return nil, err
	}

	c.SubscribeTopics([]string{cfg.Kafka.Topic}, nil)

	return &Consumer{kafkaConsumer: c}, nil
}

func (c *Consumer) Consume(cfg *config.Config) error {
	for {
		msg, err := c.kafkaConsumer.ReadMessage(-1)
		if err != nil {
			return err
		}

		// Handle the message
		log.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		go handleMessage(cfg, msg)
	}
}

func (c *Consumer) Close() {
	if err := c.kafkaConsumer.Close(); err != nil {
		log.Printf("Failed to close Kafka consumer: %v", err)
	} else {
		log.Println("Kafka consumer closed successfully")
	}
}
