package config

import (
	_ "github.com/joho/godotenv/autoload"
	"os"
)

type AppConfig struct {
	DBName string
	DBHost string
	DBUser string
	Port string
	DBPort string
	DBPassword string
	SessionKey string
	SSLMode string
}

func GetConfig() (AppConfig, error) {
   return AppConfig{
	   DBName:     os.Getenv("DB_NAME"),
	   DBHost:     os.Getenv("DB_HOST"),
	   DBUser:     os.Getenv("DB_USER"),
	   Port:       os.Getenv("PORT"),
	   DBPort:     os.Getenv("DB_PORT"),
	   DBPassword:  os.Getenv("DB_PASSWORD"),
	   SessionKey:  os.Getenv("SESSION_KEY"),
	   SSLMode:    os.Getenv("SSLMODE"),
   }, nil
}