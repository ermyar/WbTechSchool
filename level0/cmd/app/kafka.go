package main

import (
	"context"
	"log/slog"

	"github.com/ermyar/WbTechSchool/l0/internal/kafka"
	"github.com/jackc/pgx/v4"
)

type App struct {
	conn     *pgx.Conn
	consumer *kafka.Consumer
	log      *slog.Logger
}

func (a *App) Handle(ar []byte) error {
	a.log.Info("Received from Kafka", slog.String("MESSAGE", string(ar)))
	return nil
}

func (a *App) Stop(ctx context.Context) {
	a.conn.Close(ctx)
	a.consumer.Stop()
}

func (a *App) Close(ctx context.Context) {
	a.log.Info("Close: closing app's connections")
	a.conn.Close(ctx)
	a.consumer.Close()
}

func (a *App) Start() error {
	defer a.Close(context.Background())

	if err := a.consumer.Start(); err != nil {
		a.log.Error("App stopped with error", slog.String("error", err.Error()))
		return err
	}
	return nil
}
