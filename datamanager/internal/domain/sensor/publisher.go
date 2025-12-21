package sensor

type ReadingsPublisher interface {
	Publish(topic string, payload []byte) error
	PublishJson(topic string, payload any) error
}
