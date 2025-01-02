package utils

import (
	"errors"

	"github.com/google/uuid"
)

func ValidateId(id string) error {
	if _, err := uuid.Parse(id); err == nil {
		return errors.New("uuid invalid")
	}

	return nil
}

func CreateId() string {
	return uuid.Must(uuid.NewRandom()).String()
}
