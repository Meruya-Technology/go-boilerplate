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

	DbDriver   string
	DbName     string
	DbHost     string
	DbPort     string
	DbUsername string
	DbPassword string
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

	/// Initialize Db Envs
	cfg.DbDriver = os.Getenv("DB_DRIVER")
	cfg.DbName = os.Getenv("DB_NAME")
	cfg.DbHost = os.Getenv("DB_HOST")
	cfg.DbPort = os.Getenv("DB_PORT")
	cfg.DbUsername = os.Getenv("DB_USERNAME")
	cfg.DbPassword = os.Getenv("DB_PASSWORD")

	return cfg
}
