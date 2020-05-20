package token

import (
	"net/http"

	"github.com/cjburchell/ridemanager/common/service/data"
	"github.com/cjburchell/ridemanager/common/service/data/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
)

func GetUser(r *http.Request, dataService data.IService) (*models.User, error) {
	return dataService.GetUser(GetUserID(r))
}

func GetUserID(r *http.Request) models.AthleteId {
	decoded := context.Get(r, "decoded")
	claims := decoded.(jwt.MapClaims)
	userId := claims["user"].(string)
	return models.AthleteId(userId)
}
