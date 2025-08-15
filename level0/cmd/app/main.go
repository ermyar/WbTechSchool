package main

import (
	"context"
	"log/slog"
	"os"

	help "github.com/ermyar/WbTechSchool/l0/internal/pgxhelp"
)

func main() {

	// init & setup slog.Logger
	log := setupLogger()

	// Postgres Connection
	ctx := context.Background()
	conn := help.MustGetAlivePostgresConn(log, ctx)

	defer conn.Close(context.Background())
}

func setupLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
}
