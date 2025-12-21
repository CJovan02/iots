package config

import (
	"fmt"
	"os"
)

type Config struct {
	DatabaseUrl  string
	MqttBroker   string
	MqttClientId string
}

func LoadConfig() (*Config, error) {
	cfg := &Config{
		DatabaseUrl:  os.Getenv("POSTGRES_SMOKE_CONNECTION_STRING"),
		MqttBroker:   os.Getenv("MQTT_BROKER"),
		MqttClientId: os.Getenv("MQTT_CLIENT_ID"),
	}

	if cfg.DatabaseUrl == "" {
		return nil, fmt.Errorf("no postgres connection string provided")
	}
	if cfg.MqttBroker == "" {
		return nil, fmt.Errorf("MQTT broker not provided")
	}
	if cfg.MqttClientId == "" {
		fmt.Println("MQTT client id not provided, using the default 'data-manager'")
		cfg.MqttClientId = "data-manager"
	}

	return cfg, nil
}
