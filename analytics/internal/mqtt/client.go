package mqtt

import (
	json2 "encoding/json"
	"log"

	"github.com/cjovan02/iots/analytics/internal/analyser"
	"github.com/cjovan02/iots/analytics/internal/domain"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type ReadingsClient struct {
	client   mqtt.Client
	analyser *analyser.Analyser
}

// NewReadingsClient creates readings client instance and connects to broker
func NewReadingsClient(broker string, clientId string, analyser *analyser.Analyser) (*ReadingsClient, error) {
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

	return &ReadingsClient{client: client, analyser: analyser}, nil
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

	// Convert message to domain model
	var reading domain.Reading
	err := json2.Unmarshal(message.Payload(), &reading)
	if err != nil {
		log.Printf(
			"❌ error trying to unmarshal message. topic=%s, payload=%s, err=%v\n",
			message.Topic(), message.Payload(), err,
		)
		return
	}

	// Call MLAAS REST API to analyze the data
	resp, analysed, err := c.analyser.AddReading(reading)
	if err != nil {
		log.Printf("error adding reading. error=%v\n", err)
	}

	if analysed {
		log.Printf("prediction for device id: %s, is %f", reading.DeviceId, resp.Prediction)
	}
}
