package configs

import (
	"github.com/caarlos0/env/v6"
	_ "github.com/joho/godotenv/autoload" // load .env file automatically
)

var instance Config

func Load() *Config {
	if err := env.Parse(&instance); err != nil {
		panic(err)
	}

	return &instance
}

type Config struct {
	AppName    string `env:"APP_NAME"`
	AppVersion string `env:"APP_VERSION"`

	Server   Server
	Logger   Logger
	Postgres Postgres
	
}

type (
	Server struct {
		Environment       string `env:"SERVER_ENVIRONMENT"`
		Port              uint16 `env:"ADMIN_PORT"`
		MaxConnectionIdle uint16 `env:"SERVER_MAX_CONNECTION_IDLE"`
		Timeout           uint16 `env:"SERVER_TIMEOUT"`
		Time              uint16 `env:"SERVER_TIME"`
		MaxConnectionAge  uint16 `env:"SERVER_MAX_CONNECTION_AGE"`
	}
	Logger struct {
		Level    string `env:"LOGGER_LEVEL"`
		Encoding string `env:"LOGGER_ENCODING"`
	}

	Postgres struct {
		Port     uint16 `env:"POSTGRES_PORT"`
		Host     string `env:"POSTGRES_HOST"`
		Password string `env:"POSTGRES_PASSWORD"`
		User     string `env:"POSTGRES_USER"`
		Database string `env:"POSTGRES_DB"`
		PoolMax  int    `env:"POSTGRES_POOL_MAX"`
	}
)
