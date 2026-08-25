package service

import (
	"errors"
	"strings"
)

const minPasswordLength = 8

func validateEmail(email string) error {
	email = strings.TrimSpace(email)

	if email == "" {
		return errors.New("email is required")
	}

	if !strings.Contains(email, "@") {
		return errors.New("invalid email")
	}

	return nil
}

func validatePassword(password string) error {
	if password == "" {
		return errors.New("password is required")
	}

	if len(password) < minPasswordLength {
		return errors.New("password must be at least 8 characters")
	}

	return nil
}
