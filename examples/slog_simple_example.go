package main

import (
	"log/slog"
	"os"

	"github.com/libgox/slogsimple"
)

func main() {
	slog.New(slogsimple.NewHandler(&slogsimple.Config{
		Output:   os.Stdout,
		MinLevel: slog.LevelInfo,
	}))
	slog.Info("hello, slog simple")
}
