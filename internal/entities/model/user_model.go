package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;primary_key;unique;not_null"`
	Name     string    `gorm:"type:varchar(255);not_null"`
	Email    string    `gorm:"type:varchar(255);not_null;unique"`
	Password string    `gorm:"not_null"`
}
