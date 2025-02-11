package service

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWT struct {
	Base
}

func NewJWT() *JWT {
	return &JWT{}
}

func (j *JWT) GenerateToken(userId string, username string) (string, error) {
	claims := jwt.MapClaims{
		"user_id":  userId,
		"username": username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jToken, err := token.SignedString([]byte("this is secret key"))
	if err != nil {
		fmt.Println("error token")
		return "", err
	}

	return jToken, nil
}
