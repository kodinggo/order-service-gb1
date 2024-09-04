package messaging

import (
	"context"
	"log"

	"github.com/nats-io/nats.go"
)

type JetStreamRepository interface {
	AddStream(ctx context.Context, streamName string, subjects []string) error
	Publish(ctx context.Context, subject string, data []byte) error
	Subscribe(ctx context.Context, subject string, handler nats.MsgHandler) error
	ConsumeStream(ctx context.Context, streamName string, consumerName string, handler nats.MsgHandler) error
}

type jetStreamRepository struct {
	js nats.JetStreamContext
}

func NewJetStreamRepository(nc *nats.Conn) (JetStreamRepository, error) {
	js, err := nc.JetStream()
	if err != nil {
		return nil, err
	}

	return &jetStreamRepository{js: js}, nil
}

func (r *jetStreamRepository) AddStream(ctx context.Context, streamName string, subjects []string) error {
	_, err := r.js.AddStream(&nats.StreamConfig{
		Name:     streamName,
		Subjects: subjects,
		NoAck:    false,
	}, nats.Context(context.Background()))
	if err != nil {
		log.Fatalf("failed to create stream: %v", err)
	}

	return err
}

func (r *jetStreamRepository) Publish(ctx context.Context, subject string, data []byte) error {
	_, err := r.js.Publish(subject, data)
	return err
}

func (r *jetStreamRepository) Subscribe(ctx context.Context, subject string, handler nats.MsgHandler) error {
	_, err := r.js.Subscribe(subject, handler)
	return err
}

func (r *jetStreamRepository) ConsumeStream(ctx context.Context, streamName string, consumerName string, handler nats.MsgHandler) error {
	_, err := r.js.Subscribe(consumerName, handler, nats.Bind(streamName, consumerName))
	return err
}
