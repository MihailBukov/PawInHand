package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/lpernett/godotenv"
)

func NewApplicationConfiguration(envfile string) (*ApplicationConfiguration, error) {
	if envfile != "" {
		if err := godotenv.Load(envfile); err != nil {
			return nil, err
		}
	}

	result := &ApplicationConfiguration{}

	return result, envconfig.Process("", result)
}

type ApplicationConfiguration struct {
	LogLevel string `envconfig:"LOG_LEVEL" default:"debug"`

	HTTPPort               string        `envconfig:"HTTP_PORT" default:"8080"`
	HTTPServerReadTimeout  time.Duration `envconfig:"HTTP_SERVER_READ_TIMEOUT" default:"5s"`
	HTTPServerWriteTimeout time.Duration `envconfig:"HTTP_SERVER_WRITE_TIMEOUT" default:"5s"`

	DBHost                 string `envconfig:"DB_HOST" default:"localhost"`
	DBPort                 string `envconfig:"DB_PORT" default:"5432"`
	DBName                 string `envconfig:"DB_NAME" default:"database"`
	DBUsername             string `envconfig:"DB_USERNAME" default:"database"`
	DBPassword             string `envconfig:"DB_PASSWORD" default:"database"`
	DatabaseMigrationsPath string `envconfig:"DATABASE_MIGRATIONS_PATH" default:"./resources/migrations"`
	SSLMode                string `envconfig:"SSL_MODE" default:"disable"`
	BinaryParameter        string `envconfig:"BINARY_PARAMETER" default:"yes"`
}
