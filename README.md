# Internet of Things and Services (IoTS)
---
Group of projects developed for the university subject **_Internet of Things and Services_**.

The goal of those projects were to explore how the IoT systems work, through finding the appropriate dataset online in order to simulate sensor readings and sending that data do the database through servers that are implemented as microservices and are deployed in containerized environment.

## Project I - Data Management
---

This project is focused on builidng **data management** microservice system that consists of two services: ```Gateway``` and ```Datamanager```.
It also contains tool for simulating iot sensor readings, ```Sensor generator```, that reads data from _CSV_ file in batches and sends it to _Gateway_


### Data Manager
It's goal is to provide abstraction layer for accessing **PostgreSQL** database, providing CRUD and aggregation operations for manipulating data.

It's built in ```Golang``` using ```Pgx``` driver for PostgreSql. Reason for this is speed and efficiency when it comes to manipulaing data.

It works as **gRPC API** server, enabiling _Gateway_ to efficiently access stored information in the database.
Proto buffer specification is located in ```/datamanager/proto/reading.proto```


### Gateway
The name says it, it's a bridge/gateway between data and clients that want to access it.

It's built in ```.NET framework``` as a **REST API** service, exposing endpoints for communication between clients with _Datamanager_ service.
It communicates with data manager in order to access and modify data in database through **gRPC**


### Sensor generator
It simulates IoT sensor data generation and sends it to **Gateway** service.

It works by reading data from _CSV_ file in batches, sends that data to gateway, sleeps for some time and repeats the cycle.

It's built in ```Python``` as a CLI tool and it has a bunch of parameters for configuring data generation simulation.
 
Dataset (smoke detection) - https://www.kaggle.com/datasets/deepcontractor/smoke-detection-dataset/data
