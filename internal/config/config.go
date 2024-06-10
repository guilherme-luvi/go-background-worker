package config

import (
	"encoding/json"
	"os"
)

type KafkaConfig struct {
	Brokers string `json:"brokers"`
	GroupID string `json:"groupID"`
	Topic   string `json:"topic"`
}

type RedisConfig struct {
	Addr string `json:"addr"`
}

type Config struct {
	Kafka KafkaConfig `json:"kafka"`
	Redis RedisConfig `json:"redis"`
}

func LoadConfig() (*Config, error) {
	file, err := os.Open("../../config.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := &Config{}
	err = decoder.Decode(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
