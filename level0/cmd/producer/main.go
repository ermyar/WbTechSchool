package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	k "github.com/ermyar/WbTechSchool/l0/internal/kafka"
	"github.com/ermyar/WbTechSchool/l0/internal/utils"
)

const (
	kafkaAdrr  = "localhost:9091"
	kafkaTopic = "orders"
)

func main() {
	log := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	pr := k.NewProducer([]string{kafkaAdrr}, kafkaTopic, log)
	defer pr.Close()

	// creating topic if it's no there
	_, err := k.GetConn(context.Background(), kafkaAdrr, kafkaTopic)

	if err != nil {
		log.Error("can't connect with leader", slog.String("address", kafkaAdrr))
		os.Exit(1)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	if err := pr.Start(ctx, 2*time.Second); err != nil {
		log.Error("Start: finished with", utils.SlogError(err))
		return
	}
}
