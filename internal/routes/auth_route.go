package routes

import (
	authHandler "backend-template-go/internal/handler"
	tokenRepo "backend-template-go/internal/repository/token"
	userRepo "backend-template-go/internal/repository/user"
	authService "backend-template-go/internal/service/auth"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterAuthRoutes(app *fiber.App, db *gorm.DB) {
	userRepo := userRepo.NewUserRepository(db)
	tokenRepo := tokenRepo.NewTokenRepository(db)
	authService := authService.NewAuthService(userRepo, tokenRepo)
	authHandler := authHandler.NewAuthHandler(authService)

	authGroup := app.Group("/api/v1/auth")

	authGroup.Post("/register", authHandler.Register)
	authGroup.Post("/login", authHandler.Login)
	authGroup.Post("/refresh", authHandler.Refresh)
}
