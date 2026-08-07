package main

import (
	"log/slog"

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
}
