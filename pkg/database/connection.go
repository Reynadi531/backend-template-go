package database

import (
	"backend-template-go/config"
	"fmt"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func generateDSN(config config.AppConfig) string {
	if config.DB.DSN != "" {
		return config.DB.DSN
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", config.DB.Host, config.DB.Port, config.DB.User, config.DB.Password, config.DB.Name, config.DB.SSLMode)
	return dsn
}

var DB *gorm.DB

func InitDatabase() *gorm.DB {
	dsn := generateDSN(config.Config)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error().Err(err).Msg("failed connect to database")
		return nil
	}

	DB = db
	log.Info().Msg("connected to database")

	return DB
}
