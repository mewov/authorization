package storage

import (
	"database/sql"
	"log/slog"
	"time"

	_ "github.com/jackc/pgx/v5"
	"github.com/mewov/authorization/internal/config"
)

type (
	Postgres struct {
		db *sql.DB
	}
)

func ConnectToPostgres(conf *config.Config) (*Postgres, error) {
	db, err := sql.Open("postgresql", conf.PostgresUrl)
	if err != nil {
		return nil, err
	}

	for range 10 {
		err = db.Ping()
		if err == nil {
			break
		}
		slog.Info("ping database...")
		time.Sleep(time.Second * 2)
	}
	if err != nil {
		return nil, err
	}

	return &Postgres{db: db}, nil
}

func (p *Postgres) Ping() error {
	return p.db.Ping()
}

func (p *Postgres) Close() error {
	return p.db.Close()
}
