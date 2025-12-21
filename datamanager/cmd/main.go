package main

import (
	"log"
	"net"

	"github.com/CJovan02/iots/datamanager/internal/config"
	"github.com/CJovan02/iots/datamanager/internal/db"
	"github.com/CJovan02/iots/datamanager/internal/domain/sensor"
	"github.com/CJovan02/iots/datamanager/internal/grpchand"
	"github.com/CJovan02/iots/datamanager/internal/interceptor"
	"github.com/CJovan02/iots/datamanager/internal/sensorrepo"
	"github.com/CJovan02/iots/datamanager/internal/sensorsvc"
	"github.com/CJovan02/iots/datamanager/protogen/golang/sensorpg"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	mqtt "github.com/CJovan02/iots/datamanager/internal/mqtt"
)

func main() {
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
	publisher, err = mqtt.NewReadingsClient(cfg.MqttBroker, cfg.MqttClientId)
	if err != nil {
		log.Fatalf("❌ failed to connect to MQTT broker: %v", err)
	}

	// Create repo and service
	var repo sensor.Repository = sensorrepo.New(pool)
	var service sensor.Service = sensorsvc.New(repo, publisher, cfg.MqttTopic)

	// Create gRPC handler
	var sensorHandler = grpchand.NewSensorHandler(service)

	// Start server
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	// Create gRPC server
	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			interceptor.UnaryServerLoggingInterceptor,
			interceptor.UnaryServerErrMappingInterceptor,
		),
	)
	// Register service handler to server
	sensorpg.RegisterReadingsServer(server, sensorHandler)
	reflection.Register(server)

	// Start listening to requests
	log.Printf("server listening at %v", listener.Addr())
	if err = server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
