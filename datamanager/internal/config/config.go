package config

import (
	"fmt"
	"log"
	"os"
)

type Config struct {
	DatabaseUrl  string
	MqttBroker   string
	MqttClientId string
	MqttTopic    string
}

func LoadConfig() (*Config, error) {
	cfg := &Config{
		DatabaseUrl:  os.Getenv("POSTGRES_SMOKE_CONNECTION_STRING"),
		MqttBroker:   os.Getenv("MQTT_BROKER"),
		MqttClientId: os.Getenv("MQTT_CLIENT_ID"),
		MqttTopic:    os.Getenv("MQTT_TOPIC"),
	}

	if cfg.DatabaseUrl == "" {
		return nil, fmt.Errorf("POSTGRES_SMOKE_CONNECTION_STRING environment variable not set")
	}
	if cfg.MqttBroker == "" {
		return nil, fmt.Errorf("MQTT_BROKER environment variable not set")
	}
	if cfg.MqttTopic == "" {
		return nil, fmt.Errorf("MQTT_TOPIC environment variable not set")
	}
	if cfg.MqttClientId == "" {
		log.Printf("MQTT_TOPIC environment variable not set, using the default 'data-manager'")
		cfg.MqttClientId = "data-manager"
	}

	return cfg, nil
}
