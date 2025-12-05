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
	//err = sensorRepo.Create(ctx, &sensor.SensorReading{
	//	Timestamp:   time.Now(),
	//	Temperature: 0,
	//	Humidity:    0,
	//	TVOC:        0,
	//	ECO2:        0,
	//	RawHw:       0,
	//	RawEthanol:  0,
	//	PM25:        0,
	//	FireAlarm:   0,
	//})
	//if err != nil {
	//	log.Fatal(err)
	//}

	//err = sensorRepo.Update(ctx, 1, &sensor.SensorReading{
	//	Timestamp:   time.Now(),
	//	Temperature: 0,
	//	Humidity:    0,
	//	TVOC:        0,
	//	ECO2:        0,
	//	RawHw:       0,
	//	RawEthanol:  0,
	//	PM25:        0,
	//	FireAlarm:   0,
	//})
	//if err != nil {
	//	log.Fatal(err)
	//}

	err = sensorRepo.Delete(ctx, 5)
	if err != nil {
		log.Fatal(err)
	}

	//reading, err := sensorRepo.GetById(ctx, 1)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("%+v\n", *reading)

	readings, err := sensorRepo.List(ctx)
	if err != nil {
		log.Fatal(err)
	}

	for _, reading := range readings {
		fmt.Printf("%+v\n", reading)
	}
}
