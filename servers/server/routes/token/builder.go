package token

import (
	"time"

	"github.com/cjburchell/ridemanager/common/service/data/models"
	"github.com/dgrijalva/jwt-go"
)

type Builder interface {
	BuildToken(userId models.AthleteId) (string, error)
}

func GetBuilder(jwt string) Builder {
	return validator{jwt}
}

func (v validator) BuildToken(userId models.AthleteId) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": userId,
		"exp":  expirationTime,
	})

	tokenString, err := token.SignedString([]byte(v.jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
