package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type ConfigHandler struct{}

func (cfg ConfigHandler) LoadConfig() Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	/// Initialize Int Envs
	readTimeout, _ := strconv.Atoi(os.Getenv("READ_TIMEOUT"))
	WriteTimeout, _ := strconv.Atoi(os.Getenv("WRITE_TIMEOUT"))

	/// Initialize Db Envs
	dbPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	dbDriver := os.Getenv("DB_DRIVER")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	sslMode := os.Getenv("SSL_MODE")
	dbConfig := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUsername, dbPassword, dbName)

	result := Config{
		Host:         os.Getenv("HOST"),
		Port:         os.Getenv("PORT"),
		ReadTimeout:  time.Duration(readTimeout) * time.Second,
		WriteTimeout: time.Duration(WriteTimeout) * time.Second,
		DbDriver:     dbDriver,
		DbHost:       dbHost,
		DbPort:       dbPort,
		DbUsername:   dbUsername,
		DbPassword:   dbPassword,
		SslMode:      sslMode,
		DbConfig:     dbConfig,
	}

	return result
}

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
