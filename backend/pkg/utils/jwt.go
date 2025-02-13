package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var AUTH_EXPIRY time.Duration = time.Hour * 24

func GenerateJWT(subject string) (string, error) {
	key := []byte(os.Getenv("JWT_SECRET"))

	emptyToken := jwt.New(jwt.SigningMethodHS256)
	emptyToken.Claims = &jwt.RegisteredClaims{
		Issuer:    "infosphere",
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(AUTH_EXPIRY)),
		Subject:   subject,
	}

	signedToken, err := emptyToken.SignedString(key)
	if err != nil {
		fmt.Printf("error in token signature: %s\n", err.Error())
		return "", err
	}

	return signedToken, nil
}
