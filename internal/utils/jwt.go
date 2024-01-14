package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(id int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":         id,
		"authorized": true,
		"expiry":     time.Now().Add(3 * time.Minute),
	}) // TODO: this is wrong; there are specific keywords (search "jwt map claims")

	secretKey := []byte(os.Getenv("JWT_SECRET_KEY"))
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func IsValidJWT(tokenString string) error {
	_, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("signing methods do not match")
		}
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	// TODO: check token expiry and if id matches
	return err
}
