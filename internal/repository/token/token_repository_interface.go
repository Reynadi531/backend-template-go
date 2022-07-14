package token

import (
	"backend-template-go/internal/entities/model"
	"github.com/google/uuid"
)

type TokenRepository interface {
	Create(userID uuid.UUID, token string) error
	GetActiveToken(userID uuid.UUID) (model.Token, error)
	Revoke(userID uuid.UUID) error
}
