package config

import (
	"fmt"

	"github.com/caarlos0/env/v9"
)

// Config ... For connection to postgressql and server
type Config struct {
	Database Database
	Server   Server
}
type Database struct {
	Host string `env:"PGHOST" envDefault:"127.0.0.1"`
	Port string `env:"PGPORT" envDefault:"5432"`
	User string `env:"PGUSER" envDefault:"postgres"`
	Pass string `env:"PGPASSWORD" envDefault:"password"`
	Name string `env:"PGDATABASE" envDefault:"students"`
}
type Server struct {
	Port string `env:"SERVER_PORT" envDefault:"8080"`
}

func ReadConfig() (*Config, error) {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {

		return nil, fmt.Errorf("env, parse: %w", err)
	}

	return &cfg, nil
}
