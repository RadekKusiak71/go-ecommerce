package authentication

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

var secret = os.Getenv("JWT_SECERT")

func CreateJWT(accountID int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"AccountID": accountID,
		"ExpiresAt": 15000,
	})
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateJWT(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		log.Fatal(err)
		return "", nil
	}
	if !token.Valid {
		return "", fmt.Errorf("token is invalid")
	}
	return "token is valid", nil
}
