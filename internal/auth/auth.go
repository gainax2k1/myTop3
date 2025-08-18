package auth

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	// func GenerateFromPassword(password []byte, cost int) ([]byte, error)
	// bcrypt.DefaultCost is the default cost, otherwise needs to be specified
	// default is 10, 20 would be high, 5 would be low (but faster)
	// it acts as the size of exponent for hashing.
	if len(password) == 0 {
		return "", fmt.Errorf("error hashing password of zero length ")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("error hashing pass: %w", err)
	}

	return string(hashedPassword), nil
}
