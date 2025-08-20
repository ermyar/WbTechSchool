package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	k "github.com/ermyar/WbTechSchool/l0/internal/kafka"
)

func main() {
	log := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	pr := k.NewProducer([]string{"localhost:9091"}, "orders", log)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	cnt := 0
	for {
		select {
		case <-ctx.Done():
			pr.Stop()
			log.Info("STOP")
			return
		default:
			pr.Produce(context.Background(), []byte("hello"), []byte(fmt.Sprintf("produced to Kafka: %d", cnt)))
			cnt++

			time.Sleep(2 * time.Second)
		}
	}
}
