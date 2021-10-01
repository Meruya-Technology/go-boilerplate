package config

import (
	"fmt"
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
	DbPort     int
	DbUsername string
	DbPassword string
	SslMode    string
	DbConfig   string
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
	dbPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	cfg.DbDriver = os.Getenv("DB_DRIVER")
	cfg.DbName = os.Getenv("DB_NAME")
	cfg.DbHost = os.Getenv("DB_HOST")
	cfg.DbPort = dbPort
	cfg.DbUsername = os.Getenv("DB_USERNAME")
	cfg.DbPassword = os.Getenv("DB_PASSWORD")
	cfg.SslMode = os.Getenv("SSL_MODE")

	cfg.DbConfig = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DbHost, cfg.DbPort, cfg.DbUsername, cfg.DbPassword, cfg.DbName)

	return cfg
}
