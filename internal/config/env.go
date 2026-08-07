package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

const (
	DEFAULT_ADDRESS       = "0.0.0.0"
	DEFAULT_REST_PORT     = "8001"
	DEFAULT_GRPC_PORT     = "9001"
	DEFAULT_WRITE_TIMEOUT = "15"
	DEFAULT_READ_TIMEOUT  = "15"
)

func FromEnv() (*Config, error) {
	address := get("ADDRESS", DEFAULT_ADDRESS)
	restPort := get("REST_PORT", DEFAULT_REST_PORT)
	grpcPort := get("GRPC_PORT", DEFAULT_GRPC_PORT)

	writeTimeout, err := strconv.Atoi(get("WRITE_TIMEOUT", DEFAULT_WRITE_TIMEOUT))
	if err != nil {
		return nil, fmt.Errorf("$WRITE_TIMEOUT is not integer")
	}
	readTimeout, err := strconv.Atoi(get("READ_TIMEOUT", DEFAULT_READ_TIMEOUT))
	if err != nil {
		return nil, fmt.Errorf("$READ_TIMEOUT is not integer")
	}

	return &Config{
		Address:      address,
		RestPort:     restPort,
		GrpcPort:     grpcPort,
		WriteTimeout: time.Duration(writeTimeout) * time.Second,
		ReadTimeout:  time.Duration(readTimeout) * time.Second,
	}, nil
}

func get(v string, d string) string {
	k := os.Getenv(v)
	if k == "" {
		return d
	}
	return k
}
