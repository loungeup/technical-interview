package main

import (
	"context"
	"fmt"

	"github.com/loungeup/go-loungeup/pkg/log"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

func main() {
	log.Default().Debug("Starting...")

	nc, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		panic(fmt.Errorf("could not connect to NATS: %w", err))
	}
	defer nc.Close()

	js, err := jetstream.New(nc)
	if err != nil {
		panic(fmt.Errorf("could not create JetStream instance: %w", err))
	}

	stream, err := js.CreateOrUpdateStream(context.Background(), jetstream.StreamConfig{
		Name:     "test-stream",
		Subjects: []string{"test.>"},
	})
	if err != nil {
		panic(fmt.Errorf("could not create or update JetStream stream: %w", err))
	}

	consumer, err := stream.CreateOrUpdateConsumer(context.Background(), jetstream.ConsumerConfig{
		Durable:       "test-consumer",
		FilterSubject: "test.>",
	})
	if err != nil {
		panic(fmt.Errorf("could not create or update JetStream consumer: %w", err))
	}

	// Use the consumer to pull messages from the stream. For each message, log the subject.
	//
	// References:
	// 	- https://docs.nats.io/nats-concepts/jetstream/consumers#dispatch-type-pull-push
	// 	- https://github.com/nats-io/nats.go/tree/main/jetstream#continuous-polling
	_ = consumer

	log.Default().Debug("Exiting...")
}
