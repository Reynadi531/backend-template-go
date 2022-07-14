package handler

import (
	"backend-template-go/internal/entities/web"
	authService "backend-template-go/internal/service/auth"
	"backend-template-go/internal/validations"
	"github.com/gofiber/fiber/v2"
)

type authHandler struct {
	authService authService.AuthService
}

func (a authHandler) Register(c *fiber.Ctx) error {
	register := validations.RegisterAuthValidation{}
	if err := c.BodyParser(&register); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.Response{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Invalid request body",
			Data:       nil,
			Error:      "Cannot parse request body",
		})
	}

	isNotValid, err := validations.UniversalValidation(register)
	if !isNotValid {
		return c.Status(fiber.StatusBadRequest).JSON(web.Response{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Invalid request body",
			Data:       nil,
			Error:      err,
		})
	}

	res := a.authService.Register(register.Name, register.Email, register.Password)
	return c.Status(res.StatusCode).JSON(res)
}

func (a authHandler) Login(c *fiber.Ctx) error {
	login := validations.LoginAuthValidation{}
	if err := c.BodyParser(&login); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.Response{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Invalid request body",
			Data:       nil,
			Error:      "Cannot parse request body",
		})
	}

	isNotValid, err := validations.UniversalValidation(login)
	if !isNotValid {
		return c.Status(fiber.StatusBadRequest).JSON(web.Response{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Invalid request body",
			Data:       nil,
			Error:      err,
		})
	}

	result := a.authService.Login(login.Email, login.Password)
	return c.Status(result.StatusCode).JSON(result)
}

func (u authHandler) Refresh(c *fiber.Ctx) error {
	refresh := validations.RefreshAuthValidation{}
	if err := c.BodyParser(&refresh); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.Response{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Invalid request body",
			Data:       nil,
			Error:      "Cannot parse request body",
		})
	}

	isNotValid, err := validations.UniversalValidation(refresh)
	if !isNotValid {
		return c.Status(fiber.StatusBadRequest).JSON(web.Response{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Invalid request body",
			Data:       nil,
			Error:      err,
		})
	}

	result := u.authService.Refresh(refresh.RefreshToken, refresh.UserID)
	return c.Status(result.StatusCode).JSON(result)
}

func NewAuthHandler(authService authService.AuthService) AuthHandler {
	return &authHandler{authService: authService}
}
