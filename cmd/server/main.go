package main

import (
	"log/slog"
	"os"

	"go-echo-template/internal"
)

func main() {
	cfg, err := internal.LoadConfig()
	if err != nil {
		slog.Error("Could not load config", "err", err)
		os.Exit(1)
	}
	if err := internal.Run(cfg); err != nil {
		slog.Error("Failed to run server", "err", err)
		os.Exit(1)
	}
}
