package config

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	Sslmode  string
}

func NewConfig(ctx context.Context) (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, fmt.Errorf("loading .env file: %w", err)
	}
	cfg := Config{
		Host:     os.Getenv("HOST"),
		Port:     os.Getenv("PORT"),
		User:     os.Getenv("USER"),
		Password: os.Getenv("PASSWORD"),
		DBName:   os.Getenv("DBNAME"),
		Sslmode:  os.Getenv("SSLMODE"),
	}
	if cfg.Host == "" {
		return nil, errors.New("host not set")
	}
	if cfg.Port == "" {
		return nil, errors.New("port not set")
	}
	if cfg.User == "" {
		return nil, errors.New("user not set")
	}
	if cfg.Password == "" {
		return nil, errors.New("password not set")
	}
	if cfg.DBName == "" {
		return nil, errors.New("dbname not set")
	}
	if cfg.Sslmode == "" {
		return nil, errors.New("sslmode not set")
	}
	return &cfg, nil
}
