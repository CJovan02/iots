import csv
from itertools import islice

from batchCreateRequest import BatchCreateRequest
from reading import Reading

dataPath = "data/smoke_detection_iot.csv"

with open(dataPath, mode='r') as file:
    csvFile = csv.DictReader(file)
    readings: list[Reading] = []

    for row in islice(csvFile, 0, 10):
        reading = Reading.from_dict(row)
        print(reading)
        readings.append(reading)

    request = BatchCreateRequest(readings)

