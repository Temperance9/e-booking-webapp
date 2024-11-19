package utils

import (
	"errors"
	"regexp"
)

// ValidateEmail ensures the email is in a valid format.
func ValidateEmail(email string) error {
	re := regexp.MustCompile(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`)
	if !re.MatchString(email) {
		return errors.New("invalid email format")
	}
	return nil
}

// ValidatePassword ensures a password meets criteria (e.g., length, complexity).
func ValidatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}
	return nil
}
