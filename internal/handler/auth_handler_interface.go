package handler

import "github.com/gofiber/fiber/v2"

type AuthHandler interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	Refresh(c *fiber.Ctx) error
}
