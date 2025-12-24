import logging
import os

def get_gateway_url_from_dotenv() -> str:
    url = os.getenv("GATEWAY_URL")
    if url is None:
        logging.getLogger().error("Env variable 'GATEWAY_URL' not found, using default value: 'http://localhost:7002'")
        exit(1)

    return url