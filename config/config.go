package config

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"os"
)

type AppConfig struct {
	App struct {
		Name      string
		Port      string
		JWTSecret string
	}
	DB struct {
		Host     string
		Port     string
		User     string
		Password string
		Name     string
		SSLMode  string
		DSN      string
	}
}

var Config AppConfig

func InitConfig(DevMode bool) *AppConfig {
	if DevMode {
		if err := godotenv.Load(); err != nil {
			log.Error().Err(err).Msg("Error loading .env file")
		}
	}

	Config.App.Name = os.Getenv("APP_NAME")
	Config.App.Port = os.Getenv("PORT")
	Config.App.JWTSecret = os.Getenv("JWT_SECRET")

	Config.DB.Host = os.Getenv("DB_HOST")
	Config.DB.Port = os.Getenv("DB_PORT")
	Config.DB.User = os.Getenv("DB_USER")
	Config.DB.Password = os.Getenv("DB_PASSWORD")
	Config.DB.Name = os.Getenv("DB_NAME")
	Config.DB.SSLMode = os.Getenv("DB_SSL_MODE")
	Config.DB.DSN = os.Getenv("DB_DSN")

	return &Config
}
