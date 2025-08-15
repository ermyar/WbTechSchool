package pgxhelp

import (
	"context"
	"log/slog"
	"net/url"
	"os"

	"github.com/jackc/pgx/v4"
)

func getPostgresUrl(log *slog.Logger) string {
	const (
		hostAddress = "postgres:5432"
		sslsetup    = "sslmode=disable"
	)
	url := url.URL{
		Scheme: "postgres",

		User: url.UserPassword(
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASS"),
		),

		Host: hostAddress,

		Path: os.Getenv("POSTGRES_NAME"),

		RawQuery: sslsetup,
	}

	urlPG := url.String()

	log.Info("try to connect", slog.String("address", urlPG))

	return urlPG
}

func GetPostgresConn(log *slog.Logger, ctx context.Context) (*pgx.Conn, error) {

	conn, err := pgx.Connect(ctx, getPostgresUrl(log))

	return conn, err
}

func MustGetAlivePostgresConn(log *slog.Logger, ctx context.Context) *pgx.Conn {
	conn, err := GetPostgresConn(log, ctx)

	if err != nil {
		log.Error("unable to connect to Postgres", slog.String("error", err.Error()))
		os.Exit(1)
	}

	if err := conn.Ping(ctx); err != nil {
		log.Error("cant ping Postgres", slog.String("error", err.Error()))
		os.Exit(1)
	}

	log.Info("connected to Postgres succesfully")

	return conn
}
