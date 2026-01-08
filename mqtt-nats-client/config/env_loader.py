import os
import logging

def load_envs() -> tuple[str, int, str]:
    address = os.getenv("MQTT_BROKER_ADDRESS")
    if address is None:
        logging.getLogger().warning("MQTT_BROKER_ADDRESS env variable not found, using the default value 'localhost'")
        address = "localhost"

    port = os.getenv("MQTT_BROKER_PORT")
    if port is None:
        logging.getLogger().warning("MQTT_BROKER_PORT env variable not found, using the default value '1883'")
        port = 1883
    
    port = int(port)

    nats_broker = os.getenv("NATS_BROKER")
    if nats_broker is None:
        logging.getLogger().warning("NATS_BROKER env variable not found, using the default value 'nats://nats:4222'")
        nats_broker = "nats://nats:4222"


    return address, port, nats_broker
