from dotenv import load_dotenv
import logging
import os

def get_gateway_url_from_dotenv() -> str:
    load_dotenv("../docker/.env")
    port = os.getenv("GATEWAY_PORT")
    if port is None:
        logging.getLogger().error("Env variable 'GATEWAY_PORT' is not set inside ../docker/.env' file")
        exit(1)

    return "http://localhost:" + port