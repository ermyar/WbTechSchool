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
	help "github.com/ermyar/WbTechSchool/l0/internal/pgxhelp"
)

var app App

func main() {

	// init & setup slog.Logger
	log := setupLogger()

	// Postgres Connection
	ctx := context.Background()
	conn := help.MustGetAlivePostgresConn(log, ctx)

	capacity, _ := strconv.Atoi(os.Getenv("LRU_CAPACITY"))

	app = App{
		conn: conn,
		log:  log,
		lru:  lru.NewLru[string](capacity),
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
