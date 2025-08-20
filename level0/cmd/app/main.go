package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"strings"
	"syscall"

	k "github.com/ermyar/WbTechSchool/l0/internal/kafka"
	help "github.com/ermyar/WbTechSchool/l0/internal/pgxhelp"
)

func main() {

	// init & setup slog.Logger
	log := setupLogger()

	// Postgres Connection
	ctx := context.Background()
	conn := help.MustGetAlivePostgresConn(log, ctx)

	appa := App{
		conn: conn,
		log:  log,
	}

	intr, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	appa.consumer = k.NewConsumer(strings.Fields(os.Getenv("KAFKA_BROKERS")),
		os.Getenv("KAFKA_CONSUME_TOPIC"), "consumer-group-id", log, &appa, intr)

	if err := appa.Start(); err != nil {
		log.Info("Exit")
		os.Exit(1)
	}
}

func setupLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
}
