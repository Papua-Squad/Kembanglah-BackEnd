package config

import (
	"github.com/golobby/dotenv"
	"os"
)

type Config struct {
	Server struct {
		Port string `env:"SERVER_PORT"`
	}
	Database struct {
		Name     string `env:"DATABASE_NAME"`
		Host     string `env:"DATABASE_HOST"`
		Port     string `env:"DATABASE_PORT"`
		Username string `env:"DATABASE_USERNAME"`
		Password string `env:"DATABASE_PASSWORD"`
	}
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
