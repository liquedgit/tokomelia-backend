package service

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

func GenerateToken(username string, role string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["username"] = username
	claims["expired"] = time.Now().Add(time.Hour * 24).Unix()
	claims["issued"] = time.Now()
	claims["role"] = role
	strToken, err := token.SignedString("SECRET_KEY")
	if err != nil {
		log.Fatal("Error Generating Key")
		return "", err
	}
	return strToken, err
}

func ParseToken(strToken string) (string, error) {
	token, err := jwt.Parse(strToken, func(token *jwt.Token) (interface{}, error) {
		return "SECRET_KEY", nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username := claims["username"].(string)
		return username, nil
	} else {
		return "", err
	}
}
