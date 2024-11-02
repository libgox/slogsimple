package main

import (
	"os"

	"github.com/libgox/slogsimple"
	"golang.org/x/exp/slog"
)

func main() {
	slog.New(slogsimple.NewHandler(&slogsimple.Config{
		Output:   os.Stdout,
		MinLevel: slog.LevelInfo,
	}))
	slog.Info("hello, slog simple")
}
