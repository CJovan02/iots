package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	MqttBroker         string
	MqttClientID       string
	MqttSubscribeTopic string
	MlaasUrl           string
	WindowSize         uint
	NatsBroker         string
	NatsPublishSubject string
}

func LoadConfig() (*Config, error) {
	cfg := &Config{
		MqttBroker:         os.Getenv("MQTT_BROKER"),
		MqttClientID:       os.Getenv("MQTT_CLIENT_ID"),
		MqttSubscribeTopic: os.Getenv("MQTT_SUBSCRIBE_TOPIC"),
		MlaasUrl:           os.Getenv("MLAAS_URL"),
		NatsBroker:         os.Getenv("NATS_BROKER"),
		NatsPublishSubject: os.Getenv("NATS_PUBLISH_SUBJECT"),
	}

	size, err := getEnvUint("ANALYSE_WINDOW_SIZE")
	if err != nil {
		return nil, err
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

	if cfg.MlaasUrl == "" {
		return nil, fmt.Errorf("MLAAS_URL env variable not set")
	}

	if cfg.NatsBroker == "" {
		return nil, fmt.Errorf("NATS_BROKER env variable not set")
	}

	if cfg.NatsPublishSubject == "" {
		return nil, fmt.Errorf("NATS_PUBLISH_TOPIC env variable not set")
	}

	cfg.WindowSize = size

	return cfg, nil
}

func getEnvUint(key string) (uint, error) {
	val := os.Getenv(key)
	if val == "" {
		return 0, fmt.Errorf("%s env variable not set", key)
	}

	i, err := strconv.ParseUint(val, 10, 32)
	if err != nil {
		return 0, err
	}

	return uint(i), nil
}
