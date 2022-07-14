package auth

import (
	"backend-template-go/internal/entities/model"
	"backend-template-go/internal/entities/web"
	tokenRepo "backend-template-go/internal/repository/token"
	userRepo "backend-template-go/internal/repository/user"
	"backend-template-go/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type authService struct {
	userRepo  userRepo.UserRepository
	tokenRepo tokenRepo.TokenRepository
}

func (a authService) Register(name string, email string, password string) web.Response {
	userExists, err := a.userRepo.UserExists(email)
	if err != nil && err != gorm.ErrRecordNotFound {
		return web.Response{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Internal server error",
			Data:       nil,
			Error:      "Error when checking if user exists",
		}
	}

	if userExists {
		return web.Response{
			StatusCode: fiber.StatusBadRequest,
			Message:    "User already exists",
			Data:       nil,
			Error:      "User already exists",
		}
	}

	hashPassword, err := utils.GeneratePassword(password)

	if err != nil {
		return web.Response{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Internal server error",
			Data:       nil,
			Error:      "Error when hashing password",
		}
	}

	user, err := a.userRepo.Create(&model.User{
		ID:       utils.GenerateUUID(),
		Name:     name,
		Email:    email,
		Password: hashPassword,
	})

	if err != nil {
		return web.Response{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Internal server error",
			Data:       nil,
			Error:      "Error when creating user",
		}
	}

	return web.Response{
		StatusCode: fiber.StatusCreated,
		Message:    "User created",
		Data: map[string]string{
			"id":    user.ID.String(),
			"name":  user.Name,
			"email": user.Email,
		},
		Error: nil,
	}
}

func (a authService) Login(email string, password string) web.Response {
	userExist, err := a.userRepo.UserExists(email)
	if err != nil && err != gorm.ErrRecordNotFound {
		return web.Response{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Internal server error",
			Data:       nil,
			Error:      "Error when checking if user exists",
		}
	}

	if !userExist {
		return web.Response{
			StatusCode: fiber.StatusBadRequest,
			Message:    "User does not exist",
			Data:       nil,
			Error:      "User does not exist",
		}
	}

	user, err := a.userRepo.FindByEmail(email)
	if err != nil {
		return web.Response{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Internal server error",
			Data:       nil,
			Error:      "Error when finding user",
		}
	}

	if !utils.ComparePassword(user.Password, password) {
		return web.Response{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Wrong credentials",
			Data:       nil,
			Error:      "Wrong credentials",
		}
	}

	var rt string
	jwttoken, exp, err := utils.GenerateJWTToken(*user)
	if err != nil {
		return web.Response{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Internal server error",
			Data:       nil,
			Error:      "Error when generating JWT token",
		}
	}

	token, err := a.tokenRepo.GetActiveToken(user.ID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return web.Response{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Internal server error",
			Data:       nil,
			Error:      "Error when getting active token",
		}
	}

	if token.Revoked == false {
		if err := a.tokenRepo.Revoke(user.ID); err != nil {
			return web.Response{
				StatusCode: fiber.StatusInternalServerError,
				Message:    "Internal server error",
				Data:       nil,
				Error:      "Error when revoking token",
			}
		}
	}

	rt, _ = utils.GenerateRefreshToken()
	err = a.tokenRepo.Create(user.ID, rt)
	if err != nil {
		return web.Response{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Internal server error",
			Data:       nil,
			Error:      "Error when creating refresh token",
		}
	}

	return web.Response{
		StatusCode: fiber.StatusOK,
		Message:    "Login successful",
		Data: map[string]interface{}{
			"user_id":       user.ID.String(),
			"token":         jwttoken,
			"refresh_token": rt,
			"exp":           exp,
		},
	}
}

func (a authService) Refresh(token string, userId string) web.Response {
	rt, err := a.tokenRepo.GetActiveToken(uuid.MustParse(userId))
	if err != nil && err != gorm.ErrRecordNotFound {
		return web.Response{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Internal server error",
			Data:       nil,
			Error:      "Error when getting active token",
		}
	}

	if rt.Token != token {
		return web.Response{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Invalid token",
			Data:       nil,
			Error:      "Invalid token",
		}
	}

	if err := a.tokenRepo.Revoke(uuid.MustParse(userId)); err != nil {
		return web.Response{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Internal server error",
			Data:       nil,
			Error:      "Error when revoking token",
		}
	}

	newrt, _ := utils.GenerateRefreshToken()
	if err := a.tokenRepo.Create(uuid.MustParse(userId), newrt); err != nil {
		return web.Response{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Internal server error",
			Data:       nil,
			Error:      "Error when creating refresh token",
		}
	}

	user, err := a.userRepo.FindByID(userId)
	if err != nil {
		return web.Response{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Internal server error",
			Data:       nil,
			Error:      "Error when finding user",
		}
	}

	jwttoken, exp, err := utils.GenerateJWTToken(*user)
	if err != nil {
		return web.Response{
			StatusCode: fiber.StatusInternalServerError,
			Message:    "Internal server error",
			Data:       nil,
			Error:      "Error when generating JWT token",
		}
	}

	return web.Response{
		StatusCode: fiber.StatusOK,
		Message:    "Refresh successful",
		Data: map[string]interface{}{
			"user_id":       user.ID.String(),
			"token":         jwttoken,
			"refresh_token": newrt,
			"exp":           exp,
		},
	}
}

func NewAuthService(userRepo userRepo.UserRepository, tokenRepo tokenRepo.TokenRepository) AuthService {
	return &authService{userRepo: userRepo, tokenRepo: tokenRepo}
}
