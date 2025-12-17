import csv
from itertools import islice
from time import sleep

import requests

from api.batch_create_request import BatchCreateRequest
from config.parser_config import configure_parser, validate_args
from domain.reading import Reading


def batch_create(readings: list[Reading], dry_run: bool) -> None:
    if dry_run:
        print(f"Would send request with {len(readings)} readings")
        return

    print(f"Sending request with {len(readings)} readings")
    url = "http://localhost:8081/readings/batch"

    dto = BatchCreateRequest(readings=readings)

    resp = requests.post(url, json=dto.to_dict())
    if not resp.ok:
        print(resp.status_code, resp.text)
        exit(1)

    print(f"Successfully created {len(readings)} readings")


def sleep_with_message(delay: int) -> None:
    print(f"Waiting {delay} seconds...")
    sleep(delay)


def main():
    parser = configure_parser()
    args = parser.parse_args()
    validate_args(args, parser)

    file = args.file
    start = args.start
    end = args.end
    delay = args.delay
    print_readings = args.print_readings
    dry_run = args.dry_run
    batch_size = args.batch_size

    if dry_run:
        print("dry run")
    print("reading from: ", file)
    print("starting at row: ", start)
    print("ending at: ", end)
    print("delay: ", delay)
    print()

    with open(file, mode='r') as file:
        csv_file = csv.DictReader(file)
        readings: list[Reading] = []
        first_batch = True

        print("begin reading...")
        for row in islice(csv_file, start, end):
            reading = Reading.from_dict(row)
            if print_readings:
                print(row)
            readings.append(reading)

            if len(readings) >= batch_size:
                if not first_batch:
                    sleep_with_message(delay)

                print()
                batch_create(readings, dry_run)
                readings = []
                first_batch = False

        if len(readings) > 0:
            if not first_batch:
                sleep_with_message(delay)

            print()
            batch_create(readings, dry_run)


if __name__ == "__main__":
    main()
