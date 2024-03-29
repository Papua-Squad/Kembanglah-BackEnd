package config

import (
	"os"

	"github.com/golobby/dotenv"
)

type Config struct {
	Server struct {
		Port string `env:"SERVER_PORT"`
	}
	Database struct {
		Host     string `env:"DATABASE_HOST"`
		Port     string `env:"DATABASE_PORT"`
		Postgres struct {
			DB       string `env:"POSTGRES_DB"`
			User     string `env:"POSTGRES_USER"`
			Password string `env:"POSTGRES_PASSWORD"`
		}
	}
	JwtSecret string `env:"JWT_SECRET"`
}

func NewConfig() (*Config, error) {
	config := Config{}
	file, err := os.Open(".env")
	if err != nil {
		return nil, err
	}

	err = dotenv.NewDecoder(file).Decode(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
