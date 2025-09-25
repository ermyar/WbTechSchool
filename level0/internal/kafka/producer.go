package kafka

import (
	"context"
	"log/slog"
	"time"

	"github.com/ermyar/WbTechSchool/l0/internal/json"
	"github.com/ermyar/WbTechSchool/l0/internal/utils"
	"github.com/segmentio/kafka-go"
)

type Producer struct {
	w    *kafka.Writer
	log  *slog.Logger
	stop chan struct{}
}

func NewProducer(brokers []string, topic string, log *slog.Logger) *Producer {
	return &Producer{
		stop: make(chan struct{}),
		log:  log,
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

func (p *Producer) Start(ctx context.Context, t time.Duration) error {
	p.log.Info("start producing")

	for {
		select {
		case <-ctx.Done():
			p.log.Info("stopped by ctx")
			return nil
		case <-p.stop:
			p.log.Info("Stop called")
			return nil
		default:
			ord, err := json.GetRandomOrder()
			if err != nil {
				p.log.Error("error while generating ", utils.SlogError(err))
				continue
			}
			bytes, err := json.GetBytes(p.log, ord)
			if err != nil {
				p.log.Error("error while marshalling ", utils.SlogError(err))
				continue
			}
			p.Produce(ctx, []byte("order"), bytes)

			// test setting (imitation of real work)
			time.Sleep(t)
		}
	}
}

// stop producing, hand cancelling.
// runs in other goroutine, blocking func on p.stop chan
func (p *Producer) Stop() {
	p.log.Info("Stop writing")

	p.stop <- struct{}{}
}

// closing connection with Kafka
func (p *Producer) Close() {
	p.log.Info("Closing producer's connection to Kafka")
	if err := p.w.Close(); err != nil {
		p.log.Error("Kafka: Close writer", slog.String("error:", err.Error()))
	}
}

// use to connect with leader (and creates a topic when auto.create.topics.enable='true')
func GetConn(ctx context.Context, address, topic string) (*kafka.Conn, error) {
	conn, err := kafka.DialLeader(context.Background(), "tcp", address, topic, 0)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
