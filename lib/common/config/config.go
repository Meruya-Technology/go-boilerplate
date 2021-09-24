package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Host         string
	Port         string
	ReadTimeout  time.Duration /// In Seconds
	WriteTimeout time.Duration /// In Seconds
}

func (cfg Config) LoadConfig() Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	/// Initialize String Envs
	cfg.Host = os.Getenv("HOST")
	cfg.Port = os.Getenv("PORT")

	/// Initialize Int Envs
	readTimeout, _ := strconv.Atoi(os.Getenv("READ_TIMEOUT"))
	WriteTimeout, _ := strconv.Atoi(os.Getenv("WRITE_TIMEOUT"))
	cfg.ReadTimeout = time.Duration(readTimeout) * time.Second
	cfg.WriteTimeout = time.Duration(WriteTimeout) * time.Second

	return cfg
}
