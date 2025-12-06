package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/CJovan02/iots/project1-microservices/datamanager/internal/config"
	"github.com/CJovan02/iots/project1-microservices/datamanager/internal/db"
	"github.com/CJovan02/iots/project1-microservices/datamanager/internal/domain/sensor"
	"github.com/CJovan02/iots/project1-microservices/datamanager/internal/sensorrepo"
	"github.com/CJovan02/iots/project1-microservices/datamanager/internal/sensorsvc"
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

	var ctx = context.Background()
	var repo sensor.Repository = sensorrepo.New(pool)
	var service sensor.Service = sensorsvc.New(repo)

	readings, err := service.List(ctx)
	if err != nil {
		log.Fatal(err)
	}

	for _, reading := range readings {
		fmt.Printf("%+v\n", reading)
	}

	date := time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)
	st, err := service.GetStatistics(ctx, date, time.Now())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", *st)

}
