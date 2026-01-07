package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/cjovan02/iots/analytics/internal/analyser"
	"github.com/cjovan02/iots/analytics/internal/config"
	"github.com/cjovan02/iots/analytics/internal/mqtt"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	an := analyser.NewAnalyser(cfg.MlaasUrl, cfg.WindowSize)

	client, err := mqtt.NewReadingsClient(ctx, cfg.MqttBroker, cfg.MqttClientID, an)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Subscribe(cfg.MqttSubscribeTopic)
	if err != nil {
		log.Fatal(err)
	}

	<-ctx.Done()
	log.Println("shutting down...")
}
