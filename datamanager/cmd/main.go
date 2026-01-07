package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/CJovan02/iots/datamanager/internal/api"
	"github.com/CJovan02/iots/datamanager/internal/config"
	"github.com/CJovan02/iots/datamanager/internal/db"
	"github.com/CJovan02/iots/datamanager/internal/domain/sensor"
	"github.com/CJovan02/iots/datamanager/internal/grpchand"
	"github.com/CJovan02/iots/datamanager/internal/interceptor"
	mqtt "github.com/CJovan02/iots/datamanager/internal/mqtt"
	"github.com/CJovan02/iots/datamanager/internal/sensorrepo"
	"github.com/CJovan02/iots/datamanager/internal/sensorsvc"
)

func main() {
	// Unblocks ctx.Done() channel when os closes the program
	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	// Load config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Connect to db
	pool, err := db.NewPostgresPool(cfg.DatabaseUrl)
	if err != nil {
		log.Fatalf("❌ failed to connect to database: %v", err)
	}
	defer pool.Close()

	// Create client and connect to MQTT broker
	var publisher sensor.ReadingsPublisher
	publisher, err = mqtt.NewReadingsClient(ctx, cfg.MqttBroker, cfg.MqttClientId)
	if err != nil {
		log.Fatalf("❌ failed to connect to MQTT broker: %v", err)
	}

	// Create repo and service
	var repo sensor.Repository = sensorrepo.New(pool)
	var service sensor.Service = sensorsvc.New(repo, publisher, cfg.MqttTopic)

	// Create gRPC handler
	var sensorHandler = grpchand.NewSensorHandler(service)

	// Create gRPC server
	server, err := api.NewGrpcServer(
		":8080",
		sensorHandler,
		interceptor.UnaryServerLoggingInterceptor,
		interceptor.UnaryServerErrMappingInterceptor,
	)
	if err != nil {
		log.Fatal(err)
	}

	// Start the server but don't block the main thread
	go func() {
		if err := server.Run(ctx); err != nil {
			log.Printf("❌ gRPC server error: %v\n", err)
			stop()
		}
	}()

	<-ctx.Done() // channel waits for signal (os.Interrupt or syscall.SIGTERM)

	// All abstraction layers know how to close their connections to external services
	// when the context cancels
	log.Println("shutting down...")

	// Main created the pool, main will close the pool
	pool.Close()
}
