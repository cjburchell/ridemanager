package routes

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/cjburchell/ridemanager/service/data"

	"github.com/cjburchell/ridemanager/service/data/models"
	"github.com/cjburchell/ridemanager/settings"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
)

func buildToken(userId models.AthleteId) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": userId,
		"exp":  expirationTime,
	})

	tokenString, err := token.SignedString([]byte(settings.JwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func getUser(r *http.Request, dataService data.IService) (*models.User, error) {
	decoded := context.Get(r, "decoded")
	claims := decoded.(jwt.MapClaims)
	userId := claims["user"].(string)
	return dataService.GetUser(models.AthleteId(userId))
}

func validateTokenMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authorizationHeader := req.Header.Get("Authorization")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ")
			if len(bearerToken) == 2 {
				token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
					}
					return []byte(settings.JwtSecret), nil
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
