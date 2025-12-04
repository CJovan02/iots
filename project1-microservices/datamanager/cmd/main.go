package main

import (
	"fmt"
	"log"
	"time"

	"github.com/CJovan02/iots/project1-microservices/datamanager/protogen/golang/reading"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	var orderRequest = reading.GetReadingResponse{
		Timestamp:   timestamppb.New(time.Now()),
		Temperature: 30.3,
		Humidity:    20.1,
		Tvoc:        5,
		ECo2:        6,
		RawHw:       6,
		RawEthanol:  6,
		Pm_25:       5,
		FireAlarm:   0,
	}

	bytes, err := protojson.Marshal(&orderRequest)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bytes))
	fmt.Println("test")
}
