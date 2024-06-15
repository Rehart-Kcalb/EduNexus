package utils

import "errors"

func ValidatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("Пароль слишком короткий, длина должна быть >= 8")
	}
	if len(password) > 16 {
		return errors.New("Пароль слишком длинный, длина должна быть <= 16")
	}
	if HasNonPatternChars(password, `(?m)[^!-\/:-@[-`+"`{-~a-z\\d]") {
		return errors.New("Пароль должен содержать только латиницу, цифры и специальные символы")
	}

	return nil
}
