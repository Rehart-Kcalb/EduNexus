package utils

import "errors"

func ValidatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("Password too short, length >= 8")
	}
	if len(password) > 16 {
		return errors.New("Password too long, length <= 16")
	}
	if HasNonPatternChars(password, `(?m)[^!-\/:-@[-`+"`{-~a-z\\d]") {
		return errors.New("Password must contain only latin, digits and special symbols")
	}

	return nil
}
