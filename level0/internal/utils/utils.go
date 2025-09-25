package utils

import (
	"errors"
	"log/slog"
)

var ErrWrongData = errors.New("wrong data incame")

func SlogError(err error) slog.Attr {
	return slog.String("error", err.Error())
}
