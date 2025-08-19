package kafka

import (
	"context"
	"log/slog"
	"sync"

	"github.com/segmentio/kafka-go"
)

type Handler interface {
	Handle([]byte) error
}

type Consumer struct {
	mu         sync.Mutex
	stop       bool
	reader     *kafka.Reader
	handlerMsg Handler
	log        *slog.Logger
}

// Create and initiate new Consumer
func NewConsumer(address []string, topic, groupID string, log *slog.Logger, handler Handler) *Consumer {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  address,
		Topic:    topic,
		MaxBytes: 1e6, // 10Mb
		Logger:   slog.NewLogLogger(log.Handler(), slog.LevelInfo),
		// GroupID:           groupID,
		// HeartbeatInterval: 5000, // ms
	})

	return &Consumer{reader: r, stop: false, handlerMsg: handler, log: log}
}

// Start reading from Kafka
func (c *Consumer) Start(log *slog.Logger) error {
	log.Info("Kafka: start Consumer")

	for {
		c.mu.Lock()
		if c.stop {
			c.mu.Unlock()
			break
		}
		c.mu.Unlock()

		kafkaMsg, err := c.reader.ReadMessage(context.Background())
		if err != nil {
			log.Error("Kafka: while reading", slog.String("error", err.Error()))
			return err
		}

		c.handlerMsg.Handle(kafkaMsg.Value)
	}

	log.Info("Kafka: Consumer stopped")
	return nil
}

// Closing connection and stop reading from Kafka
func (c *Consumer) Stop() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.log.Info("Kafka: closing consuming")

	c.stop = true
	c.reader.Close()
}
