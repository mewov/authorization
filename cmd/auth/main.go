package main

import (
	"context"
	"log/slog"
	"time"

	"github.com/mewov/authorization/internal/config"
	"github.com/mewov/authorization/internal/storage"
)

func main() {
	conf, err := config.FromEnv()
	if err != nil {
		slog.Error("create config", "error", err)
		return
	}

	pq, err := storage.ConnectToPostgres(conf.PostgresUrl)
	if err != nil {
		slog.Error("connect to postgres", "error", err)
		return
	}
	defer pq.Close()

	ctx, done := context.WithTimeout(context.Background(), time.Second*5)
	if err := pq.Migration(ctx); err != nil {
		slog.Error("migration database", "error", err)
		return
	}

	slog.Info("success migration")
	done()
}
