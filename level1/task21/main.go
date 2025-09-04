package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"
)

// some kind of Adapter pattern use
// we have one logger(or databse, ikd, a lot of examples) in all code
// once upon a time we decided to change our logger
// and we dont want to change all code to this
// instead of this we can use Adapter pattern
type LogAdapter interface {
	Println(v ...any)
}

type slogLogger struct {
	*slog.Logger
}

func newSlogLogger() *slogLogger {
	return &slogLogger{slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))}
}

// custom this as we want
func (sl *slogLogger) Println(args ...any) {
	var ar []any
	for i, val := range args {
		ar = append(ar, slog.Any(fmt.Sprintf("msg#%d", i+1), val))
	}
	sl.Info("(log):", ar...)
}

func main() {
	var logger LogAdapter

	{
		logger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
		logger.Println("here we work")

		logger.Println("k-boom")
	}

	// we need to change our logger immediatly
	// so, we can change not a lot to do this
	// all thanks to adapter

	{
		logger = newSlogLogger()
		logger.Println("wow, it really work")

		logger.Println("this pattern do magic")
	}
}
