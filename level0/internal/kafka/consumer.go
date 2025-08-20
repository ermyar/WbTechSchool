package kafka

import (
	"context"
	"io"
	"log/slog"

	"github.com/segmentio/kafka-go"
)

type Handler interface {
	Handle([]byte) error
}

type Consumer struct {
	ctx        context.Context
	stop       chan struct{}
	reader     *kafka.Reader
	handlerMsg Handler
	log        *slog.Logger
}

// Create and initiate new Consumer.
func NewConsumer(address []string, topic, groupID string, log *slog.Logger, handler Handler, ctx context.Context) *Consumer {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  address,
		Topic:    topic,
		GroupID:  groupID,
		MaxBytes: 1e6, // 10Mb
		Logger:   slog.NewLogLogger(log.Handler(), slog.LevelInfo),
		// CommitInterval: 10 * time.Second, // 10s
	})

	return &Consumer{reader: r, ctx: ctx, stop: make(chan struct{}), handlerMsg: handler, log: log}
}

func (c *Consumer) ConsumeAndHandle() error {
	c.log.Info("Consume and Handle called")

	// fetching msg
	kafkaMsg, err := c.reader.FetchMessage(context.Background())

	if err == io.EOF {
		c.log.Info("Reader has been closed")
		return nil
	}

	if err != nil {
		c.log.Error("Kafka: while reading", slog.String("error", err.Error()))
		return err
	}

	// handle Msg
	// important not to loose Msg!
	if err = c.handlerMsg.Handle(kafkaMsg.Value); err != nil {
		c.log.Error("Handle crashed while consuming", slog.String("error", err.Error()))
		return err
	}

	// commiting Msg
	if err := c.reader.CommitMessages(context.Background(), kafkaMsg); err != nil {
		c.log.Error("Kafka: while commiting", slog.String("error", err.Error()))
		return err
	}

	return nil
}

// Start reading from Kafka
func (c *Consumer) Start() error {
	c.log.Info("Kafka: start Consumer")

	for {
		select {
		case <-c.stop:
			c.log.Info("Called Stop func")
			return nil
		case <-c.ctx.Done():
			c.log.Warn("Consumer stop caused by ctx")
			return nil
		default:
			if err := c.ConsumeAndHandle(); err != nil {
				return err
			}
		}
	}
}

// Closing connection and stop reading from Kafka
func (c *Consumer) Stop() {

	c.log.Info("Kafka: closing consuming")
	c.stop <- struct{}{}
	c.reader.Close()
}
