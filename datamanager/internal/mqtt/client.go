package mqtt

import (
	"encoding/json"
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

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}

	log.Printf("✅ connected to MQTT broker: %s\n", broker)

	return &ReadingsClient{client: client}, nil
}

func (c *ReadingsClient) Publish(topic string, payload []byte) error {
	token := c.client.Publish(topic, 0, false, payload)
	token.Wait()
	return token.Error()
}

func (c *ReadingsClient) PublishJson(topic string, payload any) error {
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	return c.Publish(topic, data)
}
