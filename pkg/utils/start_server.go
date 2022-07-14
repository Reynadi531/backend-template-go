package utils

import (
	"backend-template-go/config"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
)

func fiberConnURL() string {
	PORT := config.Config.App.Port
	if PORT == "" {
		PORT = "3000"
	}

	return fmt.Sprintf("0.0.0.0:%s", PORT)
}

func StartServerWithGracefulShutdown(a *fiber.App) {
	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		if err := a.Shutdown(); err != nil {
			log.Error().Err(err).Msg("Server is not shutting down")
		}

		close(idleConnsClosed)
	}()

	fiberConnURL := fiberConnURL()

	if err := a.Listen(fiberConnURL); err != nil {
		log.Error().Err(err).Msg("Server is not running")
	}

	<-idleConnsClosed
}

func StartServer(a *fiber.App) {
	fiberConnURL := fiberConnURL()

	if err := a.Listen(fiberConnURL); err != nil {
		log.Error().Err(err).Msg("Server is not running")
	}
}
