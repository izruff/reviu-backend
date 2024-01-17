package utils

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(id int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": strconv.FormatInt(id, 10),
		"exp": jwt.NewNumericDate(time.Now().Add(2 * time.Minute)),
	})

	secretKey := []byte(os.Getenv("JWT_SECRET_KEY"))
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func IsValidJWT(tokenString string) (int64, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("signing methods do not match") // TODO: error handling
		}

		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		return 0, err // if token expired: "token has invalid claims: token is expired"
	}

	if !token.Valid {
		return 0, errors.New("invalid token")
	}

	tokenSub, err := token.Claims.GetSubject()
	if err != nil {
		return 0, err
	}
	userID, err := strconv.ParseInt(tokenSub, 10, 64)
	if err != nil {
		return 0, errors.New("invalid ID field")
	}

	return userID, nil
}
