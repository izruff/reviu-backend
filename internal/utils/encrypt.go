package utils

import "golang.org/x/crypto/bcrypt"

func GetPasswordHash(plainPassword string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(plainPassword), 10)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func IsPasswordValid(passwordHash string, plainPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(plainPassword))
}
