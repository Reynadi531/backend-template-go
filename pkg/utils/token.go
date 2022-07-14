package utils

import (
	"backend-template-go/config"
	"backend-template-go/internal/entities/model"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"strings"
	"time"
)

var (
	JWT_SIGNATURE_KEY  = []byte(config.Config.App.JWTSecret)
	JWT_EXPIRE_TIME    = time.Now().Add(time.Hour * 24 * 7).Unix()
	JWT_SIGNING_METHOD = jwt.SigningMethodHS256
)

type MyClaims struct {
	jwt.StandardClaims
	ID   string `json:"id"`
	Name string `json:"name"`
}

func GenerateJWTToken(user model.User) (string, int64, error) {
	claims := MyClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    config.Config.App.Name,
			ExpiresAt: JWT_EXPIRE_TIME,
		},
		Name: user.Name,
		ID:   user.ID.String(),
	}

	token := jwt.NewWithClaims(JWT_SIGNING_METHOD, claims)
	signedToken, err := token.SignedString(JWT_SIGNATURE_KEY)
	if err != nil {
		return "", 0, err
	}

	return signedToken, JWT_EXPIRE_TIME, nil
}

func GenerateRefreshToken() (string, error) {
	id, err := gonanoid.New()
	return id, err
}

type TokenMetadata struct {
	UserID  uuid.UUID
	Expires int64
}

// ExtractTokenMetadata func to extract metadata from JWT.
func ExtractTokenMetadata(c *fiber.Ctx) (*TokenMetadata, error) {
	token, err := verifyToken(c)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userID, err := uuid.Parse(claims["id"].(string))
		if err != nil {
			return nil, err
		}

		expires := int64(claims["exp"].(float64))

		return &TokenMetadata{
			UserID:  userID,
			Expires: expires,
		}, nil
	}

	return nil, err
}

func extractToken(c *fiber.Ctx) string {
	bearToken := c.Get("Authorization")

	onlyToken := strings.Split(bearToken, " ")
	if len(onlyToken) == 2 {
		return onlyToken[1]
	}

	return ""
}

func verifyToken(c *fiber.Ctx) (*jwt.Token, error) {
	tokenString := extractToken(c)

	token, err := jwt.Parse(tokenString, jwtKeyFunc)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func jwtKeyFunc(token *jwt.Token) (interface{}, error) {
	return JWT_SIGNATURE_KEY, nil
}
