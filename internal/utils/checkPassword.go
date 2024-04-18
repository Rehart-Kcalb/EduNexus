package utils

import "golang.org/x/crypto/bcrypt"

func CheckPassword(password, hashedPassword string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(hashedPassword)); err != nil {
		return false
	}
	return true
}
