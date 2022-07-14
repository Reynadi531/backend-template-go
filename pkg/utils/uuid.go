package utils

import (
	uuid "github.com/google/uuid"
)

func GenerateUUID() uuid.UUID {
	uuid, _ := uuid.NewUUID()
	return uuid
}
