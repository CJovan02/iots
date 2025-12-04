package config

import (
	"fmt"
	"os"
)

type Config struct {
	DatabaseUrl string
}

func LoadConfig() (*Config, error) {
	cfg := &Config{
		DatabaseUrl: os.Getenv("POSTGRES_SMOKE_CONNECTION_STRING"),
	}

	if cfg.DatabaseUrl == "" {
		return nil, fmt.Errorf("no postgres connection string provided")
	}

	return cfg, nil
}
