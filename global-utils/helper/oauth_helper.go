package helper

import (
	"fmt"
	"os"
	"strings"

	jwt "github.com/golang-jwt/jwt/v5"
)

func GetAuthorizationValue(token string) string {
	authorizationToken := strings.Split(token, " ")
	return authorizationToken[1]
}

func VerifyToken(tokenString string, tokenType string) (*jwt.Token, error) {
	var secretKey string

	secretKey = os.Getenv("JWT_SECRET_KEY")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}
