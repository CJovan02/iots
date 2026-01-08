import asyncio
import signal
from logging import getLogger
from config.logger_config import configure_logging
from config.env_loader import load_config
from mqtt_client.mqtt_client import MqttClient
from nats_client.nats_client import NatsClient

logger = getLogger()

async def main():
    configure_logging(False)
    config = load_config()

    # Initialize clients
    mqtt_client = MqttClient()
    mqtt_client.connect_and_subscribe(
        config.mqtt_address, config.mqtt_port, config.mqtt_subscribe_topic
    )
    mqtt_client.loop_start()  # runs mqtt client in separate thread

    nats_client = NatsClient()
    await nats_client.connect(config.nats_broker)
    await nats_client.subscribe(config.nats_subscribe_subject)

    stop_event = asyncio.Event()

    loop = asyncio.get_running_loop()
    stop_signals = [signal.SIGINT, signal.SIGTERM]
    for sig in stop_signals:
        loop.add_signal_handler(sig, stop_event.set)

    # Blocks main thread until one of the stop signals arrive
    await stop_event.wait()

    logger.info("Exiting...")

    await nats_client.disconnect()
    mqtt_client.disconnect()


if __name__ == "__main__":
    asyncio.run(main())
