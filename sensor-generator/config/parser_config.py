import argparse


def configure_parser() -> argparse.ArgumentParser:
    parser = __generate_parser()
    __add_arguments(parser)
    return parser


def __generate_parser() -> argparse.ArgumentParser:
    return argparse.ArgumentParser(
        description=
        'Simulator for smoke detection IoT data readings.\n'
        'It reads data from .csv file and sends it to gateway API to write data in database.\n'
        'It sends data in batches, delay is used to specify delay between requests.\n\n'
        "NOTE: All of the parameters are optional, they all have default values",
        epilog='Example: python generator.py --file data.csv --start 5 --end 200 --delay 0.5',
        formatter_class=argparse.RawDescriptionHelpFormatter,
    )


def __add_arguments(parser: argparse.ArgumentParser) -> None:
    parser.add_argument(
        '--file', '-f',
        help='Path to data file, must be csv format with specific columns.\nSee data/smoke_detection_iot.csv for example, '
             'this is also the default file.',
        default='data/smoke_detection_iot.csv',
    )
    parser.add_argument('--start', '-s', type=int, default=0, help='start reading from this row (inclusive)')
    parser.add_argument('--end', '-e', type=int, default=100, help='stop reading at this row (inclusive)')
    parser.add_argument('--delay', '-d', type=float, default=1.0, help='delay between batch requests in seconds')
    parser.add_argument('--dry-run', action='store_true', help='run without sending data to gateway API')
    parser.add_argument("--version", action='version', version='%(prog)s 1.0.0')


def validate_args(args: argparse.Namespace, parser: argparse.ArgumentParser) -> None:
    if args.start > args.end:
        parser.error("start must be less than end")

    if args.start < 0 or args.end < 0:
        parser.error("start and end must be positive numbers")

    if args.delay < 0 or args.delay > 20:
        parser.error("Delay must be between 0 and 20")
