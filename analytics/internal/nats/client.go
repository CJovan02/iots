package nats

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

type PredictionsClient struct {
	conn           *nats.Conn
	publishSubject string
}

func NewPredictionsClient(ctx context.Context, broker string, publishSubject string) (*PredictionsClient, error) {
	opts := []nats.Option{
		nats.ReconnectWait(2 * time.Second),
		nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
			log.Printf("Disconnected from NATS: %v", err)
		}),
		nats.ReconnectHandler(func(nc *nats.Conn) {
			log.Printf("Reconnected to NATS at %s", nc.ConnectedUrl())
		}),
		nats.ClosedHandler(func(nc *nats.Conn) {
			log.Println("NATS connection closed")
		}),
	}

	conn, err := nats.Connect(broker, opts...)
	if err != nil {
		log.Println("error connecting to NATS")
		return nil, err
	}

	log.Printf("✅ connected to NATS at %s\n", broker)

	// Graceful stop
	go func() {
		<-ctx.Done()
		if conn.IsClosed() {
			log.Println("context canceled, closing NATS connection")
			conn.Drain()
			conn.Close()
		}
	}()

	return &PredictionsClient{
		conn:           conn,
		publishSubject: publishSubject,
	}, nil
}

func (c *PredictionsClient) Publish(payload []byte) error {
	return c.conn.Publish(c.publishSubject, payload)
}

func (c *PredictionsClient) PublishJson(payload any) error {
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	return c.Publish(data)
}
