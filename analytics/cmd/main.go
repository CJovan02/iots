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
	"github.com/cjovan02/iots/analytics/internal/nats"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	an := analyser.NewAnalyser(cfg.MlaasUrl, cfg.WindowSize)

	// Create NATS client
	natsCl, err := nats.NewPredictionsClient(ctx, cfg.NatsBroker, cfg.NatsPublishSubject)
	if err != nil {
		log.Fatal(err)
	}

	// Create MQTT client
	mqttCl, err := mqtt.NewReadingsClient(ctx, cfg.MqttBroker, cfg.MqttClientID, an, natsCl)
	if err != nil {
		log.Fatal(err)
	}

	err = mqttCl.Subscribe(cfg.MqttSubscribeTopic)
	if err != nil {
		log.Fatal(err)
	}

	<-ctx.Done()
	log.Println("shutting down...")
}
