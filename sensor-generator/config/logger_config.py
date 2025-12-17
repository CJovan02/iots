import logging

def configure_logging(verbose: bool) -> logging.Logger:
    logging.basicConfig(
        level= logging.DEBUG if verbose else logging.INFO,
        format='[%(asctime)s] [%(levelname)s] %(message)s',
        datefmt='%H:%M:%S'
    )

    return logging.getLogger(__name__)


