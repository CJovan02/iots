package config

import (
	"fmt"
	"os"
)

type Config struct {
	MqttBroker         string
	MqttClientID       string
	MqttSubscribeTopic string
}

func LoadConfig() (*Config, error) {
	cfg := &Config{
		MqttBroker:         os.Getenv("MQTT_BROKER"),
		MqttClientID:       os.Getenv("MQTT_CLIENT_ID"),
		MqttSubscribeTopic: os.Getenv("MQTT_SUBSCRIBE_TOPIC"),
	}

	if cfg.MqttBroker == "" {
		return nil, fmt.Errorf("MQTT_BROKER env variable not set")
	}

	if cfg.MqttClientID == "" {
		return nil, fmt.Errorf("MQTT_CLIENT_ID env variable not set")
	}

	if cfg.MqttSubscribeTopic == "" {
		return nil, fmt.Errorf("MQTT_SUBSCRIBE_TOPIC env variable not set")
	}

	return cfg, nil
}
