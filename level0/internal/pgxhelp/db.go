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

type Connection struct {
	conn *pgx.Conn
	log  *slog.Logger
}

func GetPostgresConn(log *slog.Logger, ctx context.Context) (*Connection, error) {

	urlPG := getPostgresUrl(log)
	log.Info("try to connect", slog.String("address", urlPG))

	conn, err := pgx.Connect(ctx, urlPG)
	return &Connection{conn: conn, log: log}, err
}

func (c *Connection) CheckAlive(ctx context.Context) {
	if err := c.conn.Ping(ctx); err != nil {
		c.log.Error("Postgres: cant ping", slog.String("error", err.Error()))
		os.Exit(1)
	}
}

func MustGetAlivePostgresConn(log *slog.Logger, ctx context.Context) *Connection {
	conn, err := GetPostgresConn(log, ctx)

	if err != nil {
		conn.log.Error("Postgres: unable to connect", slog.String("error", err.Error()))
		os.Exit(1)
	}

	conn.CheckAlive(ctx)

	conn.log.Info("Postgres: connected succesfully")

	return conn
}

func (c *Connection) Close(ctx context.Context) {
	c.log.Info("Postgres: closing connection")
	c.conn.Close(ctx)
}
