import os
import logging
from dataclasses import dataclass

@dataclass
class Config:
    mqtt_address: str
    mqtt_port: int
    mqtt_subscribe_topic: str
    nats_broker: str
    nats_subscribe_subject: str

def load_config() -> Config:
    address = os.getenv("MQTT_BROKER_ADDRESS")

    if address is None:
        logging.getLogger().warning("MQTT_BROKER_ADDRESS env variable not found, using the default value 'localhost'")
        address = "localhost"

    port = os.getenv("MQTT_BROKER_PORT")
    if port is None:
        logging.getLogger().warning("MQTT_BROKER_PORT env variable not found, using the default value '1883'")
        port = 1883

    port = int(port)

    mqtt_topic = os.getenv("MQTT_SUBSCRIBE_TOPIC")
    if mqtt_topic is None:
        logging.getLogger().error("MQTT_SUBSCRIBE_TOPIC env variable not found")
        raise RuntimeError("MQTT_SUBSCRIBE_TOPIC env variable not found")

    nats_broker = os.getenv("NATS_BROKER")
    if nats_broker is None:
        logging.getLogger().warning("NATS_BROKER env variable not found, using the default value 'nats://nats:4222'")
        nats_broker = "nats://nats:4222"

    nats_subject = os.getenv("NATS_SUBSCRIBE_SUBJECT")
    if nats_subject is None:
        logging.getLogger().error(
            "NATS_SUBSCRIBE_SUBJECT env variable not found")
        raise RuntimeError("NATS_SUBSCRIBE_SUBJECT env variable not found")


    return Config(
        mqtt_address=address,
        mqtt_port=port,
        mqtt_subscribe_topic=mqtt_topic,
        nats_broker=nats_broker,
        nats_subscribe_subject=nats_subject,
    )
