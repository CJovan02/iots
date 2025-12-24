import os
import logging

def load_envs() -> tuple[str, str]:
    address = os.getenv("MQTT_BROKER_ADDRESS")
    if address is None:
        logging.getLogger().warning("MQTT_BROKER_ADDRESS env variable not found, using the default value 'localhost'")
        address = "localhost"

    port = os.getenv("MQTT_BROKER_PORT")
    if port is None:
        logging.getLogger().warning("MQTT_BROKER_PORT env variable not found, using the default value '1883'")
        port = 1883
    
    port = int(port)

    return [address, port]
