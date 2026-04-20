package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const SECRET_KEY = "supersecret"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"userId": userId,
		"exp": time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString([]byte(SECRET_KEY))
}

func VerifyToken(token string) error {
	parsedToken, err := jwt.Parse(token, func(tk *jwt.Token) (any, error) {
		_, ok := tk.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Unexpected signing method")
		}
		return []byte(SECRET_KEY), nil
	})
	if err != nil {
		return errors.New("Could not parse token")
	}
	if !parsedToken.Valid {
		return errors.New("Invalid token")
	}
	// claims, ok := parsedToken.Claims.(jwt.MapClaims)
	// if ! ok {
	// 	return errors.New("Invalid token claims")
	// }
	// email := claims["email"].(string)
	// userId := claims["userId"].(int64)

	return nil
}
