package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	j "github.com/ermyar/WbTechSchool/l0/internal/json"
	k "github.com/ermyar/WbTechSchool/l0/internal/kafka"
)

const (
	kafkaAdrr  = "localhost:9091"
	kafkaTopic = "orders"
)

func main() {
	log := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	pr := k.NewProducer([]string{kafkaAdrr}, kafkaTopic, log)

	_, err := k.GetConn(context.Background(), kafkaAdrr, kafkaTopic)

	if err != nil {
		log.Error("can't connect with leader", slog.String("address", kafkaAdrr))
		os.Exit(1)
	}

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
			ord, err := j.GetRandomOrder()
			if err != nil {
				log.Error("error while generating ", slog.String("error", err.Error()))
				return
			}
			bytes, err := j.GetBytes(log, ord)
			if err != nil {
				log.Error("error while marshalling ", slog.String("error", err.Error()))
				return
			}
			pr.Produce(context.Background(), []byte("order"), bytes)
			cnt++

			time.Sleep(2 * time.Second)
		}
	}
}
