package utils

import (
	"errors"
	"regexp"
)

// ValidateLogin checks if the login is either a valid username or a valid email address.
func ValidateLogin(login string) error {
	// Check if the login is an email
	if isValidEmail(login) {
		return nil
	}

	// Check username conditions
	if len(login) < 8 {
		return errors.New("Логин слишком короткий, длина должна быть >= 8")
	}
	if len(login) > 32 {
		return errors.New("Логин слишком длинный, длина должна быть <= 32")
	}

	if HasNonPatternChars(login, "[^a-z]") {
		return errors.New("Должна быть только латиница")
	}

	return nil
}

// isValidEmail checks if the string is a valid email address.
func isValidEmail(email string) bool {
	// Simple email regex pattern
	const emailPattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailPattern)
	return re.MatchString(email)
}
