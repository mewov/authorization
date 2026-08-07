package config

import "time"

type (
	Config struct {
		Address      string
		RestPort     string
		GrpcPort     string
		WriteTimeout time.Duration
		ReadTimeout  time.Duration
		PostgresUrl  string
	}
)
