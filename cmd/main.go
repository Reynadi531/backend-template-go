package main

import (
	"backend-template-go/config"
	"backend-template-go/internal/routes"
	"backend-template-go/internal/validations"
	"backend-template-go/pkg/database"
	"backend-template-go/pkg/middleware"
	"backend-template-go/pkg/utils"
	"flag"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

var (
	DevMode = flag.Bool("dev", false, "dev mode")
)

func init() {
	flag.Parse()
	config.InitConfig(*DevMode)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	if *DevMode {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	database.InitDatabase()
	database.AutoMigrate(database.DB)
}

func main() {
	app := fiber.New()

	validations.InitValidations()

	middleware.RegisterMiddleware(app)

	routes.RegisterAuthRoutes(app, database.DB)

	if *DevMode {
		utils.StartServer(app)
	} else {
		utils.StartServerWithGracefulShutdown(app)
	}
}
