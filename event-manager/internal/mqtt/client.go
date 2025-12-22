package mqtt

import (
	"encoding/json"
	"log"

	"github.com/cjovan02/iots/event-manager/internal/domain"
	"github.com/cjovan02/iots/event-manager/internal/dto"
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type ReadingsClient struct {
	client       mqtt.Client
	thresholds   *domain.EventThresholds
	publishTopic string
}

// NewReadingsClient creates readings client instance and connects to broker
func NewReadingsClient(broker string, clientId string, thresholds *domain.EventThresholds, publishTopic string) (*ReadingsClient, error) {
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

	return &ReadingsClient{client: client, thresholds: thresholds, publishTopic: publishTopic}, nil
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

	var response dto.ReadingResponse
	err := json.Unmarshal(message.Payload(), &response)
	if err != nil {
		log.Printf(
			"❌ error trying to unmarshal message. topic=%s, payload=%s, err=%v\n",
			message.Topic(), message.Payload(), err,
		)
		return
	}

	event, detected, err := detectEvent(response, *c.thresholds)
	if err != nil {
		log.Printf("❌ detect event error: %v\n", err)
		return
	}

	if !detected {
		return
	}

	log.Println("smoke event detected")

	err = c.PublishJson(c.publishTopic, *event)
	if err != nil {
		log.Printf("❌ error publishing event, %v\n", err)
		return
	}
}

func detectEvent(reading dto.ReadingResponse, thresholds domain.EventThresholds) (*domain.SmokeEvent, bool, error) {
	detected := false
	var triggers []*domain.Trigger
	var event = &domain.SmokeEvent{
		ReadingId: reading.Id,
		Timestamp: reading.Timestamp,
		Triggers:  nil,
	}

	pm25 := reading.PM25
	tvoc := float64(reading.Tvoc)
	eco2 := float64(reading.ECo2)
	temp := reading.Temperature

	if pm25 > thresholds.Pm25 {
		triggers = append(triggers, domain.NewPm25Trigger(pm25, thresholds.Pm25))
		detected = true
	}
	if tvoc > thresholds.Tvoc {
		triggers = append(triggers, domain.NewTvocTrigger(tvoc, thresholds.Tvoc))
		detected = true
	}
	if eco2 > thresholds.Eco2 {
		triggers = append(triggers, domain.NewEco2Trigger(eco2, thresholds.Eco2))
		detected = true
	}
	if temp > thresholds.Temperature {
		triggers = append(triggers, domain.NewTemperatureTrigger(temp, thresholds.Temperature))
		detected = true
	}

	event.Triggers = triggers

	return event, detected, nil
}
