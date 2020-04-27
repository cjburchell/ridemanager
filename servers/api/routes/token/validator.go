package token

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
)

type Validator interface {
	ValidateMiddleware(next http.HandlerFunc) http.HandlerFunc
}

type validator struct {
	jwtSecret string
}

func GetValidator(jwt string) Validator {
	return validator{jwt}
}

func (v validator) ValidateMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authorizationHeader := req.Header.Get("Authorization")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
					}
					return []byte(v.jwtSecret), nil
				})

				if err != nil {
					w.WriteHeader(http.StatusUnauthorized)
					return
				}
				if token.Valid {
					context.Set(req, "decoded", token.Claims)
					next(w, req)
				} else {
					w.WriteHeader(http.StatusUnauthorized)
				}
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	})
}
