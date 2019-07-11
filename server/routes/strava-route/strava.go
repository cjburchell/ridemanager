package strava_route

import (
	"encoding/json"
	"github.com/cjburchell/go-uatu"
	"github.com/cjburchell/ridemanager/routes/token"
	"github.com/cjburchell/ridemanager/service/data"
	"github.com/cjburchell/ridemanager/service/strava"
	"github.com/gorilla/mux"
	"net/http"
)

func Setup(r *mux.Router, service data.IService) {
	dataRoute := r.PathPrefix("/api/v1/strava").Subrouter()

	dataRoute.HandleFunc("/segments/starred", token.ValidateMiddleware(
		func(writer http.ResponseWriter, request *http.Request) {
			handleGetStarredSegments(writer, request, service)
		})).Methods("GET")
}

func handleGetStarredSegments(writer http.ResponseWriter, request *http.Request, service data.IService) {
	user, err := token.GetUser(request, service)
	if err != nil{
		writer.WriteHeader(http.StatusUnauthorized)
		log.Error(err)
		return
	}

	stravaService := strava.NewService(user.StravaToken)

	segments, err := stravaService.StaredSegments()
	if err != nil{
		writer.WriteHeader(http.StatusBadRequest)
		log.Error(err)
		return
	}

	reply, _ := json.Marshal(segments)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	_, err = writer.Write(reply)
	if err != nil {
		log.Error(err)
	}
}
