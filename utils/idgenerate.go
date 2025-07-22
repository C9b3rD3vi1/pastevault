package utils

import (
	"github.com/google/uuid"
)

// GenerateID generates a new UUID string. to identify resources.
func GenerateID() string {
	return uuid.New().String()
}
