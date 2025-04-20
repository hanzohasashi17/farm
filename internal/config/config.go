package config

import (
	"log"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	AppPort string `env:"APP_PORT" envDefault:"8080"`

	DBHost string `env:"DB_HOST"`
	DBPort string `env:"DB_PORT"`
	DBUser string `env:"DB_USER"`
	DBPass string `env:"DB_PASSWORD"`
	DBName string `env:"DB_NAME"`
}

var cfg *Config

func GetConfig() *Config {
	if cfg != nil {
		return cfg
	}

	cfg = &Config{}
	if err := env.Parse(cfg); err != nil {
		log.Fatalf("Ошибка загрузки env: %v", err)
	}
	return cfg
}
