package database

import (
	"context"
	"encoding/json"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/guilherme-luvi/go-queue-worker/internal/config"
)

var ctx = context.Background()

type CacheClient struct {
	redisClient *redis.Client
}

func NewCacheClient(cfg *config.Config) *CacheClient {
	client := redis.NewClient(&redis.Options{
		Addr: cfg.Redis.Addr,
	})

	return &CacheClient{redisClient: client}
}

func UpdateCache(cfg *config.Config, data []byte) error {
	var record map[string]interface{}
	if err := json.Unmarshal(data, &record); err != nil {
		return err
	}

	client := NewCacheClient(cfg)
	// Assuming the record contains an ID and other fields
	id := record["id"].(string)
	err := client.redisClient.HSet(ctx, id, record).Err()
	if err != nil {
		return err
	}

	log.Printf("Cache updated for record ID: %s", id)
	return nil
}
