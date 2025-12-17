import csv
from itertools import islice

from batchCreateRequest import BatchCreateRequest
from config.parser_config import configure_parser, validate_args
from reading import Reading


def main():
    parser = configure_parser()
    args = parser.parse_args()
    validate_args(args, parser)

    file = args.file
    start = args.start
    end = args.end
    delay = args.delay
    dry_run = args.dry_run

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

        print("begin reading...")
        for row in islice(csv_file, start, end + 1):
            reading = Reading.from_dict(row)
            print(reading)
            readings.append(reading)

        print()
        if dry_run:
            print(f"Would send {len(readings)} readings")
            exit(0)

        request = BatchCreateRequest(readings)


if __name__ == "__main__":
    main()
