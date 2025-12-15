package main

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/CJovan02/iots/datamanager/internal/config"
	"github.com/CJovan02/iots/datamanager/internal/db"
	"github.com/CJovan02/iots/datamanager/internal/domain/sensor"
	"github.com/CJovan02/iots/datamanager/internal/grpchand"
	"github.com/CJovan02/iots/datamanager/internal/sensorrepo"
	"github.com/CJovan02/iots/datamanager/internal/sensorsvc"
	"github.com/CJovan02/iots/datamanager/protogen/golang/sensorpg"
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

	readings := []*sensor.Reading{
		{
			Timestamp:   time.Now().Add(-2 * time.Minute).UTC(),
			Temperature: 22.5,
			Humidity:    45.0,
			Tvoc:        120,
			ECo2:        450,
			RawHw:       13500,
			RawEthanol:  21000,
			PM25:        8.3,
			FireAlarm:   0,
		},
		{
			Timestamp:   time.Now().Add(-1 * time.Minute).UTC(),
			Temperature: 68.2,
			Humidity:    18.4,
			Tvoc:        950,
			ECo2:        2200,
			RawHw:       18000,
			RawEthanol:  16000,
			PM25:        145.7,
			FireAlarm:   1,
		},
		{
			Timestamp:   time.Now().UTC(),
			Temperature: 24.1,
			Humidity:    42.8,
			Tvoc:        140,
			ECo2:        520,
			RawHw:       14200,
			RawEthanol:  12500,
			PM25:        10.1,
			FireAlarm:   0,
		},
	}

	_, err = service.BatchCreate(context.Background(), readings)
	if err != nil {
		log.Fatal(err)
	}

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
	reflection.Register(server)

	// Start listening to requests
	log.Printf("server listening at %v", listener.Addr())
	if err = server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
