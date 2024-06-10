package queue

import (
	"context"
	"log"
	"sync"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/guilherme-luvi/go-queue-worker/pkg/config"
)

type Consumer struct {
	kafkaConsumer *kafka.Consumer
	waitGroup     sync.WaitGroup
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

func (c *Consumer) Consume(ctx context.Context, cfg *config.Config) error {
	messageChannel := make(chan *kafka.Message, cfg.NumWorkers)

	// Start worker pool
	for i := 0; i < cfg.NumWorkers; i++ {
		c.waitGroup.Add(1)
		go c.worker(ctx, cfg, messageChannel)
	}

	for {
		select {
		case <-ctx.Done():
			close(messageChannel)
			c.waitGroup.Wait()
			return nil
		default:
			msg, err := c.kafkaConsumer.ReadMessage(-1)
			if err != nil {
				return err
			}
			messageChannel <- msg
		}
	}
}

func (c *Consumer) worker(ctx context.Context, cfg *config.Config, messages <-chan *kafka.Message) {
	defer c.waitGroup.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case msg, ok := <-messages:
			if !ok {
				return
			}
			// Handle the message
			log.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
			if err := handleMessage(cfg, msg); err != nil {
				log.Printf("Error handling message: %v", err)
			}
		}
	}
}

func (c *Consumer) Close() {
	if err := c.kafkaConsumer.Close(); err != nil {
		log.Printf("Failed to close Kafka consumer: %v", err)
	} else {
		log.Println("Kafka consumer closed successfully")
	}
}
