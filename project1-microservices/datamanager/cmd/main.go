package main

import (
	"log"
	"net"

	"github.com/CJovan02/iots/project1-microservices/datamanager/internal/config"
	"github.com/CJovan02/iots/project1-microservices/datamanager/internal/db"
	"github.com/CJovan02/iots/project1-microservices/datamanager/internal/domain/sensor"
	"github.com/CJovan02/iots/project1-microservices/datamanager/internal/grpchand"
	"github.com/CJovan02/iots/project1-microservices/datamanager/internal/sensorrepo"
	"github.com/CJovan02/iots/project1-microservices/datamanager/internal/sensorsvc"
	"github.com/CJovan02/iots/project1-microservices/datamanager/protogen/golang/sensorpg"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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

	// Create repo and service
	var repo sensor.Repository = sensorrepo.New(pool)
	var service sensor.Service = sensorsvc.New(repo)

	// Create gRPC handler
	var sensorHandler = grpchand.NewSensorHandler(service)

	// Start server
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	// Create gRPC server
	server := grpc.NewServer()
	// Register service handler to server
	sensorpg.RegisterReadingsServer(server, sensorHandler)
	// Enable reflection. It is required for testing with grpcurl
	reflection.Register(server)

	// Start listening to requests
	log.Printf("server listening at %v", listener.Addr())
	if err = server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
