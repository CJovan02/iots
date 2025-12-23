import argparse


def configure_parser() -> argparse.ArgumentParser:
    parser = __generate_parser()
    __add_arguments(parser)
    return parser


def __generate_parser() -> argparse.ArgumentParser:
    return argparse.ArgumentParser(
        description=(
            "Simulator for smoke detection IoT data readings.\n"
            "Reads 'batch-size' number of rows from the provided CSV file, sends them to the API, "
            "and waits for 'delay' seconds before sending the next batch\n"
            "This simulates IoT sensor readings in real time.\n"
            "Use 'start' and 'end' to specify which rows should be sent.\n\n"
            "All parameters are optional and have default values."
        ),
        epilog='Example: python generator.py --file data.csv --start 5 --end 200 --batch-size 3 --delay 10 --url http://localhost:8080',

        formatter_class=argparse.RawDescriptionHelpFormatter,
    )


def __add_arguments(parser: argparse.ArgumentParser) -> None:
    parser.add_argument('--start', '-s', type=int, default=24950,
                        help='start reading from this row (inclusive). (default: 24950)')
    parser.add_argument('--end', '-e', type=int, default=25200,
                        help='stop reading at this row (exclusive). (default: 25200)')
    parser.add_argument('--batch-size', '-bs', type=int, default=5,
                        help='number of readings sent per API request. (default: 5)')
    parser.add_argument('--delay', '-d', type=float, default=5,
                        help='delay between batch requests in seconds. (default: 5)')
    parser.add_argument(
        '--file', '-f',
        help='Path to data file, must be csv format with specific columns.\nSee data/smoke_detection_iot.csv for example, '
             'this is also the default file.',
        default='data/smoke_detection_iot.csv',
    )
    parser.add_argument('--url', '-u', type=str,
                        help="gateway url. NOTE: if you pass this argument the default url from the root/docker/.env won't be used. "
                             "This is only used when you are not running the server with docker compose and you want to manually enter gateway url.")
    parser.add_argument('--verbose', '-ver', action='store_true',
                        help='enable debug logging (prints all readings and internal info)')
    parser.add_argument('--dry-run', '-dr', action='store_true', help='run without sending data to API')
    parser.add_argument("--version", '-v', action='version', version='%(prog)s 1.0.0')


def validate_args(args: argparse.Namespace, parser: argparse.ArgumentParser) -> None:
    if args.start > args.end:
        parser.error("start must be less than end")

    if args.start < 0 or args.end < 0:
        parser.error("start and end must be positive numbers")

    if args.delay < 0 or args.delay > 20:
        parser.error("Delay must be between 0 and 20")

    if args.batch_size < 0 or args.batch_size > 100:
        parser.error("Batch size must be between 0 and 100")
