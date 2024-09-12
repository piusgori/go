package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	result, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		return "", err
	}

	resultString := string(result)

	return resultString, nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}