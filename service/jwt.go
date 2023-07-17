package service

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/liquedgit/tokoMeLia/helper"
	"time"
)

type JwtCustomClaim struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

var jwtSecret = []byte(helper.GoDotEnvVariables("SECRET_JWT"))

func GenerateToken(ctx context.Context, username string, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &JwtCustomClaim{
		Username: username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	})

	strToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return strToken, nil
}

func JwtValidate(ctx context.Context, token string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(token, &JwtCustomClaim{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("signing method error")
		}
		return jwtSecret, nil
	})
}
