import paho.mqtt.client as mqtt
import logging
import json

def configure_mqtt_client():
    mqttc = mqtt.Client(mqtt.CallbackAPIVersion.VERSION2)
    mqttc.on_connect = __on_connect
    mqttc.on_subscribe = __on_subscribe
    mqttc.on_unsubscribe = __on_unsubscribe
    mqttc.on_message = __on_message

    return mqttc

def __on_message(client, userdata, message):
    logger = logging.getLogger()

    payload = json.loads(message.payload.decode())
    pretty = json.dumps(payload, indent=2)

    logger.info("Received message from topic: %s\n%s\n", message.topic, pretty)

def __on_subscribe(client, userdata, mid, reason_code_list, properties):
    logger = logging.getLogger()

    if reason_code_list[0].is_failure:
        logger.debug(f"Broker rejected you subscription: {reason_code_list[0]}")
    else:
        logger.debug(f"Broker granted the following QoS: {reason_code_list[0].value}")

def __on_unsubscribe(client, userdata, mid, reason_code_list, properties):
    logger = logging.getLogger()

    if len(reason_code_list) == 0 or not reason_code_list[0].is_failure:
        logger.info("unsubscribe succeeded (if SUBACK is received in MQTTv3 it success)")
    else:
        logger.info(f"Broker replied with failure: {reason_code_list[0]}")

    client.disconnect()

def __on_connect(client, userdata, flags, reason_code, properties):
    logger = logging.getLogger()

    if reason_code.is_failure:
        logger.info(f"❌ Failed to connect: {reason_code}.")
        logger.debug("loop_forever() will retry connection")
    else:
        logger.info("✅ Connected to message broker")
        logger.debug(reason_code)
        client.subscribe("event-manager/threshold-readings")