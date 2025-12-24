from config.mqtt_client_config import configure_mqtt_client
from config.logger_config import configure_logging
from config.env_loader import load_envs
import logging

def main():
    try:
        logger = configure_logging(False)
        address, port = load_envs()

        mqttc = configure_mqtt_client();
        mqttc.connect(address, port)
        mqttc.loop_forever()
    except KeyboardInterrupt:
        print()
        logging.getLogger().info("Exiting...")
        exit(0)


    
if __name__ == "__main__":
    main()
