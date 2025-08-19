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

	consm := k.NewConsumer(strings.Fields(os.Getenv("KAFKA_BROKERS")), os.Getenv("KAFKA_CONSUME_TOPIC"), "group-id", log, conn)

	intr, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// gracefull shutdown
	go func() {
		// signal came, should stop
		<-intr.Done()

		consm.Stop()
		conn.Close(context.Background())
	}()

	if err := consm.Start(log); err != nil {
		os.Exit(2)
	}
}

func setupLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
}
