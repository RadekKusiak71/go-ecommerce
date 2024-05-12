package authentication

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	psw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(psw), nil
}

func ComparePasswords(psw1, psw2 []byte) bool {
	return bcrypt.CompareHashAndPassword(psw1, psw2) != nil
}
