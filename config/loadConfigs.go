package config

import (
	"os"

	"github.com/Nahuel-Adrian-Doval/Golang-API-REST-base/models"
)

func LoadDBConfig() models.DBConfig {
	return models.DBConfig{
		DB_HOST:     os.Getenv("DB_HOST"),
		DB_PORT:     os.Getenv("DB_PORT"),
		DB_USER:     os.Getenv("DB_USER"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_NAME:     os.Getenv("DB_NAME"),
		DB_SSLMODE:  os.Getenv("DB_SSLMODE"),
		DB_TIMEZONE: os.Getenv("DB_TIMEZONE"),
	}
}

func LoadServerConfig() models.ServerConfig {
	return models.ServerConfig{
		SERVER_PORT:   os.Getenv("SERVER_PORT"),
		SERVER_PREFIX: os.Getenv("SERVER_PREFIX"),
	}
}
