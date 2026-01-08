from dataclasses import dataclass
import paho.mqtt.client as mqtt
import logging
import json

from paho.mqtt.enums import CallbackAPIVersion
from paho.mqtt.client import Client

from mqtt_client.user_data import MqttClientUserData


@dataclass
class MqttClient:
    __client: mqtt.Client | None = None

    def connect_and_subscribe(self, broker: str, port: int, topic: str):
        self.__client = mqtt.Client(callback_api_version=CallbackAPIVersion.VERSION2,
                                    userdata=MqttClientUserData(subscribe_topic=topic))
        self.__client.on_connect = self.__on_connect
        self.__client.on_subscribe = MqttClient.__on_subscribe
        self.__client.on_unsubscribe = MqttClient.__on_unsubscribe
        self.__client.on_message = MqttClient.__on_message

        self.__client.connect(broker, port)

    # Run the mqtt on main thread and block it
    def loop_forever(self):
        self.__client.loop_forever()

    # Rune the mqtt on another thread, without blocking the main thread
    def loop_start(self):
        self.__client.loop_start()

    def disconnect(self):
        if self.__client and self.__client.is_connected():
            logging.getLogger().info("Disconnecting MQTT client")
            self.__client.disconnect()

    @staticmethod
    def __on_connect(client, userdata, flags, reason_code, properties):
        logger = logging.getLogger()

        if reason_code != 0:
            logger.info(f"❌ Failed to connect: {reason_code}.")
            logger.debug("loop_forever() will retry connection")
            raise RuntimeError("MQTT connection failed")
        else:
            logger.info("✅ Connected to MQTT message broker")
            logger.debug(reason_code)
            client.subscribe(userdata.subscribe_topic)

    @staticmethod
    def __on_message(client, userdata, message):
        logger = logging.getLogger()

        payload = json.loads(message.payload.decode())
        pretty = json.dumps(payload, indent=2)

        logger.info("Received message from topic: %s\n%s\n", message.topic, pretty)

    @staticmethod
    def __on_subscribe(client: Client, userdata, mid, reason_code_list, properties):
        logger = logging.getLogger()

        if reason_code_list[0].is_failure:
            logger.debug(f"Broker rejected you subscription: {reason_code_list[0]}")
        else:
            logger.info(f"MQTT Client subscribed to topic: {userdata.subscribe_topic}")
            logger.debug(f"Broker granted the following QoS: {reason_code_list[0].value}")

    @staticmethod
    def __on_unsubscribe(client, userdata, mid, reason_code_list, properties):
        logger = logging.getLogger()

        if len(reason_code_list) == 0 or not reason_code_list[0].is_failure:
            logger.info("unsubscribe succeeded (if SUBACK is received in MQTTv3 it success)")
        else:
            logger.info(f"Broker replied with failure: {reason_code_list[0]}")

        client.disconnect()
