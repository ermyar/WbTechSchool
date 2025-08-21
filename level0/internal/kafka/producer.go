package kafka

import (
	"context"
	"log/slog"

	"github.com/segmentio/kafka-go"
)

type Producer struct {
	w   *kafka.Writer
	log *slog.Logger
}

func NewProducer(brokers []string, topic string, log *slog.Logger) *Producer {
	return &Producer{
		log: log,
		w: kafka.NewWriter(kafka.WriterConfig{
			Brokers: brokers,
			Topic:   topic,
			Logger:  slog.NewLogLogger(log.Handler(), slog.LevelInfo),
		}),
	}
}

func (p *Producer) Produce(ctx context.Context, key, msg []byte) error {
	p.log.Info("Kafka: writing msg")
	if err := p.w.WriteMessages(ctx,
		kafka.Message{
			Key:   key,
			Value: msg,
		},
	); err != nil {
		p.log.Error("Kafka: Write msg", slog.String("error:", err.Error()))
		return err
	}

	return nil
}

func (p *Producer) Stop() error {
	p.log.Info("Stop writing")

	if err := p.w.Close(); err != nil {
		p.log.Error("Kafka: Close writer", slog.String("error:", err.Error()))
		return err
	}

	return nil
}

// use to connect with leader (when auto.create.topics.enable='true')
func GetConn(ctx context.Context, address, topic string) (*kafka.Conn, error) {
	conn, err := kafka.DialLeader(context.Background(), "tcp", address, topic, 0)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
