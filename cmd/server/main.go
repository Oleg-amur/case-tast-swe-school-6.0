package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/Oleg-amur/case-task-swe-school-6.0/internal/config"
	"github.com/Oleg-amur/case-task-swe-school-6.0/internal/database"
)

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	if err := runApp(log); err != nil {
		log.Error("fatal error: %w", err)
		os.Exit(1)
	}
}

func runApp(log *slog.Logger) error {
	ctx := context.Background()

	cfg, err := config.LoadConfig("configs/config.yaml")

	if err != nil {
		return err
	}

	db, err := database.InitDb(ctx, cfg.Database.ConnectionString, log)
	if err != nil {
		return err
	}

	if err := database.RunMigrations(ctx, db, log); err != nil {
		return err
	}

	log.Info("server started", "port", cfg.Server.Port)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	log.Info("shutting down...")

	return nil
}
