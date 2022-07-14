package middleware

import (
	"backend-template-go/pkg/utils"
	"github.com/gofiber/fiber/v2"

	jwtMiddleware "github.com/gofiber/jwt/v3"
)

func JWTProtected() func(*fiber.Ctx) error {
	config := jwtMiddleware.Config{
		SigningKey:   utils.JWT_SIGNATURE_KEY,
		ContextKey:   "jwt",
		ErrorHandler: jwtError,
	}

	return jwtMiddleware.New(config)
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"error":   true,
			"details": err.Error(),
		})
	}

	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"status":  fiber.StatusBadRequest,
		"error":   true,
		"details": err.Error(),
	})
}
