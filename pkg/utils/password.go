package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func GeneratePassword(password string) (string, error) {
	bytePassword := []byte(password)

	hash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func ComparePassword(hashPassword string, inputPassword string) bool {
	byteHash := []byte(hashPassword)
	byteInput := []byte(inputPassword)

	if err := bcrypt.CompareHashAndPassword(byteHash, byteInput); err != nil {
		return false
	}

	return true
}
