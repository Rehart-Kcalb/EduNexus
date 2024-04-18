package utils

import "errors"

func ValidateLogin(login string) error {
	if len(login) < 8 {
		return errors.New("Login too short, length >= 8")
	}
	if len(login) > 32 {
		return errors.New("Login too long, length <= 32")
	}

	if HasNonPatternChars(login, "[^a-z]") {
		return errors.New("Must be only latin characters")
	}

	return nil
}
