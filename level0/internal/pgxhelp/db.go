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

	return url.String()
}

func GetPostgresConn(log *slog.Logger, ctx context.Context) (*pgx.Conn, error) {

	urlPG := getPostgresUrl(log)
	log.Info("try to connect", slog.String("address", urlPG))

	conn, err := pgx.Connect(ctx, urlPG)
	return conn, err
}

func MustGetAlivePostgresConn(log *slog.Logger, ctx context.Context) *pgx.Conn {
	conn, err := GetPostgresConn(log, ctx)

	if err != nil {
		log.Error("Postgres: unable to connect", slog.String("error", err.Error()))
		os.Exit(1)
	}

	if err := conn.Ping(ctx); err != nil {
		log.Error("Postgres: cant ping", slog.String("error", err.Error()))
		os.Exit(1)
	}

	log.Info("Postgres: connected succesfully")

	return conn
}
