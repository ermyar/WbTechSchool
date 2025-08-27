package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"

	k "github.com/ermyar/WbTechSchool/l0/internal/kafka"
	"github.com/ermyar/WbTechSchool/l0/internal/lru"
)

var app App

func main() {

	// init & setup slog.Logger
	log := setupLogger()

	capacity, err := strconv.Atoi(os.Getenv("LRU_CAPACITY"))

	if err != nil {
		log.Info("LRU_CAPACITY env is not a number")
		os.Exit(1)
	}

	app = App{
		log: log,
		lru: lru.NewLru[string](capacity),
	}

	intr, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	app.consumer = k.NewConsumer(strings.Fields(os.Getenv("KAFKA_BROKERS")),
		os.Getenv("KAFKA_CONSUME_TOPIC"), "consumer-group-id", log, &app, intr)

	if err := app.Start(); err != nil {
		log.Info("Exit")
		os.Exit(1)
	}
}

func setupLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
}
