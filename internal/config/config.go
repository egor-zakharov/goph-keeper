package config

import (
	"flag"
	"os"
)

// Config - struct
type Config struct {
	FlagLogLevel    string
	FlagDB          string
	FlagRunGRPCAddr string
}

// NewConfig - constructor Config
func NewConfig() *Config {
	return &Config{}
}

// ParseFlag - parsing command line, env flags or config file
func (c *Config) ParseFlag() {
	flag.StringVar(&c.FlagLogLevel, "l", "debug", "log level")
	flag.StringVar(&c.FlagDB, "d", "postgres://admin:admin@localhost:5432/db?sslmode=disable", "database dsn")
	flag.StringVar(&c.FlagRunGRPCAddr, "g", "localhost:8081", "address and port to run grpc server")

	flag.Parse()

	if envLogLevel := os.Getenv("LOG_LEVEL"); envLogLevel != "" {
		c.FlagLogLevel = envLogLevel
	}

	if envDB := os.Getenv("DATABASE_DSN"); envDB != "" {
		c.FlagDB = envDB
	}

	if envGRPCAddress := os.Getenv("GRPC_ADDR"); envGRPCAddress != "" {
		c.FlagRunGRPCAddr = envGRPCAddress
	}
}
