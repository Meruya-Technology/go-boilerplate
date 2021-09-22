package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Host string
	Port string
}

func (cfg Config) LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg.Host = os.Getenv("HOST")
	cfg.Port = os.Getenv("PORT")
	return cfg
}
