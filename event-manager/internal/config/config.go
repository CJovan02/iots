package config

import (
	"fmt"
	"log"
	"os"
)

type Config struct {
	Broker         string
	ClientId       string
	SubscribeTopic string
	PublishTopic   string
}

func LoadConfig() (*Config, error) {
	cfg := &Config{
		Broker:         os.Getenv("MQTT_BROKER"),
		ClientId:       os.Getenv("MQTT_CLIENT_ID"),
		SubscribeTopic: os.Getenv("MQTT_SUBSCRIBE_TOPIC"),
		PublishTopic:   os.Getenv("MQTT_PUBLISH_TOPIC"),
	}

	if cfg.Broker == "" {
		return nil, fmt.Errorf("MQTT_BROKER environment variable not set")
	}
	if cfg.SubscribeTopic == "" {
		return nil, fmt.Errorf("MQTT_SUBSCRIBE_TOPIC environment variable not set")
	}
	if cfg.PublishTopic == "" {
		return nil, fmt.Errorf("MQTT_PUBLISH_TOPIC environment variable not set")
	}
	if cfg.ClientId == "" {
		log.Printf("MQTT_CLIENT_ID environment variable not set, using the default 'event-manager'")
		cfg.ClientId = "event-manager"
	}
	return cfg, nil
}
