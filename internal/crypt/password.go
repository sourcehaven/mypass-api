package crypt

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) ([]byte, error) {
	hashedPw, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return nil, err
	}
	return hashedPw, nil
}

func IsValidPassword(plainPw string, hashedPw []byte) bool {
	err := bcrypt.CompareHashAndPassword(hashedPw, []byte(plainPw))
	return err == nil
}
