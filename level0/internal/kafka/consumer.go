package kafka

import (
	"context"
	"errors"
	"io"
	"log/slog"

	"github.com/ermyar/WbTechSchool/l0/internal/utils"

	"github.com/segmentio/kafka-go"
)

type Handler interface {
	Handle(context.Context, []byte) error
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
	})

	return &Consumer{reader: r, ctx: ctx, stop: make(chan struct{}), handlerMsg: handler, log: log}
}

func (c *Consumer) ConsumeAndHandle(ctx context.Context) error {
	c.log.Info("Consume and Handle called")

	// fetching msg
	kafkaMsg, err := c.reader.FetchMessage(ctx)

	if err == io.EOF {
		c.log.Info("Reader has been closed")
		return nil
	}

	if err != nil {
		c.log.Error("Kafka: while reading", utils.SlogError(err))
		return err
	}

	// handle Msg
	// important not to loose Msg!
	// but Msg can be wrong! in this case returns special error: WrongDataErr
	if err = c.handlerMsg.Handle(ctx, kafkaMsg.Value); err != nil {
		c.log.Error("Handle crashed while consuming", utils.SlogError(err))
		if !errors.Is(err, utils.ErrWrongData) {
			return err
		}
	}

	// commiting Msg
	if err := c.reader.CommitMessages(ctx, kafkaMsg); err != nil {
		c.log.Error("Kafka: while commiting", utils.SlogError(err))
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
			if err := c.ConsumeAndHandle(c.ctx); err != nil {
				return err
			}
		}
	}
}

// Closing connection and stop reading from Kafka.
// Blocking function to chan c.stop.
func (c *Consumer) Stop() {
	c.log.Info("Kafka: stopping consuming")
	c.stop <- struct{}{}
}

// Closing connection with Kafka.
// Should be called after Start finish.
// non-blocking func to chan c.stop.
func (c *Consumer) Close() {
	c.log.Info("Kafka: closing consumer's connection")

	// may be extra error handling
	if err := c.reader.Close(); err != nil {
		c.log.Error("Kafka: Close writer", utils.SlogError(err))
	}
}
