package service

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type JWT struct {
	Base
}

func NewJWT() *JWT {
	return &JWT{}
}

func (j *JWT) GenerateToken(userId int, username string) (string, error) {
	claims := jwt.MapClaims{
		"user_id":  userId,
		"username": username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString("this is secret key")
}
