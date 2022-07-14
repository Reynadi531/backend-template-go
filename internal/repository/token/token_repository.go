package token

import (
	"backend-template-go/internal/entities/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type tokenRepository struct {
	db *gorm.DB
}

func (t tokenRepository) Create(userID uuid.UUID, token string) error {
	return t.db.Create(&model.Token{UserID: userID, Token: token}).Error
}

func (t tokenRepository) GetActiveToken(userID uuid.UUID) (model.Token, error) {
	var token model.Token
	err := t.db.Where("user_id = ? AND revoked = ?", userID, false).First(&token).Error
	return token, err
}

func (t tokenRepository) Revoke(userID uuid.UUID) error {
	return t.db.Model(&model.Token{}).Where("user_id = ?", userID).Update("revoked", true).Error
}

func NewTokenRepository(db *gorm.DB) TokenRepository {
	return &tokenRepository{db: db}
}
