from config.mqtt_client_config import configure_mqtt_client
from config.logger_config import configure_logging
import logging

def main():
    try:
        logger = configure_logging(False)

        mqttc = configure_mqtt_client();
        mqttc.connect("localhost", 1883)
        mqttc.loop_forever()
    except KeyboardInterrupt:
        print()
        logging.getLogger().info("Exiting...")
        exit(0)


    
if __name__ == "__main__":
    main()
