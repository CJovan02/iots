package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/cjovan02/iots/event-manager/internal/config"
	"github.com/cjovan02/iots/event-manager/internal/mqtt"
)

func main() {
	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt, syscall.SIGTERM,
	)
	defer stop()

	// Load env variables
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Create client and connect to broker
	client, err := mqtt.NewReadingsClient(cfg.Broker, cfg.ClientId, cfg.Thresholds, cfg.PublishTopic)
	if err != nil {
		log.Fatal(err)
	}

	// Subscribe to provided topic
	err = client.Subscribe(cfg.SubscribeTopic)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("🚀 event manager is running")

	<-ctx.Done()

	log.Println("shutting down...")
	client.Disconnect()
}
