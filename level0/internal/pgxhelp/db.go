package pgxhelp

import (
	"context"
	"log/slog"
	"net/url"
	"os"

	"github.com/ermyar/WbTechSchool/l0/internal/utils"
	"github.com/jackc/pgx/v4"
)

func getPostgresUrl() string {
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

type PgConnection struct {
	conn *pgx.Conn
	log  *slog.Logger
}

func (pgc *PgConnection) Ping(ctx context.Context) error {
	pgc.log.Info("pinging our Postgres connection")
	return pgc.conn.Ping(ctx)
}

func (pgc *PgConnection) Close(ctx context.Context) {
	pgc.log.Info("Closing Postgres Connection")
	pgc.conn.Close(ctx)
}

func GetPostgresConn(log *slog.Logger, ctx context.Context) (*PgConnection, error) {
	urlPG := getPostgresUrl()
	log.Info("try to connect", slog.String("address", urlPG))

	conn, err := pgx.Connect(ctx, urlPG)
	return &PgConnection{conn, log}, err
}

func MustGetAlivePostgresConn(log *slog.Logger, ctx context.Context) *PgConnection {
	conn, err := GetPostgresConn(log, ctx)

	if err != nil {
		log.Error("Postgres: unable to connect", utils.SlogError(err))
		os.Exit(1)
	}

	if err := conn.Ping(ctx); err != nil {
		log.Error("Postgres: cant ping", utils.SlogError(err))
		os.Exit(1)
	}

	log.Info("Postgres: connected succesfully")

	return conn
}
