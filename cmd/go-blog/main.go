package main

import (
	"log/slog"
	"os"

	"github.com/pseudoerr/go-blog/internal/config"
	"github.com/pseudoerr/go-blog/internal/db"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)
	log.Info("starting go-blog", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")

	db, err := db.NewPostgresDB(cfg.DBURL)
	if err != nil {
		log.Error("failed to connect to db", slog.Any("error", err))
		os.Exit(1)
	}
	defer db.Close()
	log.Info("db connected")

	// Todo: init router (gin)

	// Todo: run server
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return log
}
