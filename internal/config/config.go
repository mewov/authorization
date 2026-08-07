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

func (c *Config) GetAddress() string {
	return c.Address
}

func (c *Config) GetRestPort() string {
	return c.RestPort
}

func (c *Config) GetGrpcPort() string {
	return c.GrpcPort
}

func (c *Config) GetWriteTimeout() time.Duration {
	return c.WriteTimeout
}

func (c *Config) GetReadTimeout() time.Duration {
	return c.ReadTimeout
}

func (c *Config) GetPostgresUrl() string {
	return c.PostgresUrl
}
