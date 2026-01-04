package mqtt

import (
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type ReadingsClient struct {
	client mqtt.Client
}

// NewReadingsClient creates readings client instance and connects to broker
func NewReadingsClient(broker string, clientId string) (*ReadingsClient, error) {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(broker)
	opts.ClientID = clientId
	// After 40 seconds of not sending any message, send PINGREQ to inform broker that this client is still alive
	opts.KeepAlive = 40
	opts.OnConnect = func(client mqtt.Client) {
		log.Printf("✅ connected to MQTT broker: %s\n", broker)
	}
	opts.OnConnectionLost = func(client mqtt.Client, err error) {
		log.Printf("❌ connection lost from MQTT broker: %s\n", broker)
	}

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}

	return &ReadingsClient{client: client}, nil
}

func (c *ReadingsClient) Subscribe(topic string) error {
	token := c.client.Subscribe(topic, 0, c.handleMessage)
	if token.Wait() && token.Error() != nil {
		return token.Error()
	}

	log.Printf("✅ subscribed to topic: %s\n", topic)
	return nil
}

func (c *ReadingsClient) Disconnect() {
	c.client.Disconnect(250)
}

func (c *ReadingsClient) handleMessage(_ mqtt.Client, message mqtt.Message) {
	log.Printf("received message from topic: %s\n", message.Topic())
}
