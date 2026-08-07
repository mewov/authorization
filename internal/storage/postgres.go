package storage

import (
	"context"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	pingTimeout = 5 * time.Second
	maxRetries  = 10
	retryDelay  = 2 * time.Second
)

type Postgres struct {
	db *pgxpool.Pool
}

func ConnectToPostgres(url string) (*Postgres, error) {
	ctx, cancel := context.WithTimeout(context.Background(), pingTimeout)
	defer cancel()

	pool, err := pgxpool.New(ctx, url)
	if err != nil {
		return nil, err
	}

	for i := 0; i < maxRetries; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), pingTimeout)

		err = pool.Ping(ctx)
		cancel()

		if err == nil {
			slog.Info("connected to postgres")
			return &Postgres{db: pool}, nil
		}

		slog.Warn(
			"failed to connect to postgres",
			"attempt", i+1,
			"max_attempts", maxRetries,
			"error", err,
		)

		time.Sleep(retryDelay)
	}

	pool.Close()
	return nil, err
}

func (p *Postgres) Ping() error {
	ctx, cancel := context.WithTimeout(context.Background(), pingTimeout)
	defer cancel()

	return p.db.Ping(ctx)
}

func (p *Postgres) Close() {
	if p.db != nil {
		p.db.Close()
	}
}

func (p *Postgres) Pool() *pgxpool.Pool {
	return p.db
}
