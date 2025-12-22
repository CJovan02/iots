package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/cjovan02/iots/event-manager/internal/domain"
)

type Config struct {
	Broker         string
	ClientId       string
	SubscribeTopic string
	PublishTopic   string
	Thresholds     *domain.EventThresholds
}

func LoadConfig() (*Config, error) {
	pm25, err := getEnvFloat("PM_25_THRESHOLD")
	if err != nil {
		return nil, err
	}

	tvoc, err := getEnvFloat("TVOC_THRESHOLD")
	if err != nil {
		return nil, err
	}

	eco2, err := getEnvFloat("ECO2_THRESHOLD")
	if err != nil {
		return nil, err
	}

	temp, err := getEnvFloat("TEMPERATURE_THRESHOLD")
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		Broker:         os.Getenv("MQTT_BROKER"),
		ClientId:       os.Getenv("MQTT_CLIENT_ID"),
		SubscribeTopic: os.Getenv("MQTT_SUBSCRIBE_TOPIC"),
		PublishTopic:   os.Getenv("MQTT_PUBLISH_TOPIC"),
		Thresholds: &domain.EventThresholds{
			Pm25:        pm25,
			Tvoc:        tvoc,
			Eco2:        eco2,
			Temperature: temp,
		},
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

func getEnvFloat(key string) (float64, error) {
	val := os.Getenv(key)
	if val == "" {
		return 0, fmt.Errorf("%s environment variable not set", key)
	}

	f, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid float value set for %s: %v", key, err)
	}

	return f, nil
}
