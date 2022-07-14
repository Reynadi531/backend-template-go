package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func RegisterMiddleware(app *fiber.App) {
	app.Use(
		logger.New(),
		cors.New(),
		recover.New(),
		compress.New(),
	)
}
