package main

import (
	"context"
	"fmt"
	"log"

	"github.com/CJovan02/iots/project1-microservices/datamanager/internal/config"
	"github.com/CJovan02/iots/project1-microservices/datamanager/internal/db"
	"github.com/CJovan02/iots/project1-microservices/datamanager/internal/domain/sensor"
	"github.com/CJovan02/iots/project1-microservices/datamanager/internal/repository"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	pool, err := db.NewPostgresPool(cfg.DatabaseUrl)
	if err != nil {
		log.Fatalf("❌ failed to connect to database: %v", err)
	}
	defer pool.Close()

	var sensorRepo sensor.Repository = repository.NewPgRepository(pool)

	ctx := context.Background()
	reading, err := sensorRepo.GetById(ctx, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", *reading)
}
