package mqtt

import (
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func CreateClientAndConnect(broker string, clientId string) (mqtt.Client, error) {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(broker)
	opts.ClientID = clientId

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}

	log.Printf("✅ connected to MQTT broker: %s\n", broker)
	return client, nil
}
