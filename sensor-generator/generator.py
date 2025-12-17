import csv
from itertools import islice
from time import sleep

import requests
import logging

from api.batch_create_request import BatchCreateRequest
from config.env_config import get_gateway_url_from_dotenv
from config.logger_config import configure_logging
from config.parser_config import configure_parser, validate_args
from domain.reading import Reading


def batch_create(readings: list[Reading], url: str, dry_run: bool) -> None:
    logger = logging.getLogger()
    url += "/readings/batch"
    if dry_run:
        logger.info(f"Would send request with {len(readings)} readings")
        return

    logger.info("Sending batch (%d readings)", len(readings))

    dto = BatchCreateRequest(readings=readings)

    resp = requests.post(url, json=dto.to_dict())
    if not resp.ok:
        logger.error(
            "Failed to send request",
            resp.status_code,
            resp.text,
        )
        exit(1)

    logger.info("Batch sent successfully (%d readings)", len(readings))


def sleep_with_message(delay: int) -> None:
    logging.getLogger().debug("Sleeping for %.1f seconds", delay)
    sleep(delay)


def main():
    try:
        parser = configure_parser()
        args = parser.parse_args()
        validate_args(args, parser)

        file = args.file
        start = args.start
        end = args.end
        delay = args.delay
        dry_run = args.dry_run
        batch_size = args.batch_size
        verbose = args.verbose
        arg_url = args.url

        logger = configure_logging(verbose)

        url = get_gateway_url_from_dotenv()
        if arg_url is not None:
            url = arg_url

        logger.info(
            f"Gateway URL: {url}\n\n"
            "Arguments:\n"
            f"File path: {file}\n"
            f"Start at row: {start}\n"
            f"End on row: {end}\n"
            f"Batch Size: {batch_size}\n"
            f"Delay between requests: {delay}\n"
        )

        with open(file, mode='r') as file:
            csv_file = csv.DictReader(file)
            readings: list[Reading] = []
            first_batch = True

            logger.info("Start reading from CSV file")
            for row in islice(csv_file, start, end):
                reading = Reading.from_dict(row)
                logger.debug(f"Reading: {reading}")
                readings.append(reading)

                if len(readings) >= batch_size:
                    if not first_batch:
                        sleep_with_message(delay)

                    print()
                    batch_create(readings, url, dry_run)
                    readings = []
                    first_batch = False

            if len(readings) > 0:
                if not first_batch:
                    sleep_with_message(delay)

                print()
                batch_create(readings, url, dry_run)

            print()
            logger.info("Finished. Total readings sent: %d", len(readings))
    except KeyboardInterrupt:
        print()
        logging.getLogger().info("Exiting...")
        exit(0)


if __name__ == "__main__":
    main()
