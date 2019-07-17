package user_route

import (
	"encoding/json"
	"github.com/cjburchell/ridemanager/service/data/models"
	"net/http"

	"github.com/cjburchell/go-uatu"

	"github.com/cjburchell/ridemanager/routes/token"

	"github.com/cjburchell/ridemanager/service/data"
	"github.com/gorilla/mux"
)

func Setup(r *mux.Router, service data.IService) {
	dataRoute := r.PathPrefix("/api/v1/user").Subrouter()
	dataRoute.HandleFunc("/me", token.ValidateMiddleware(
		func(writer http.ResponseWriter, request *http.Request) {
			handleGetUser(writer, request, service)
		})).Methods("GET")

	dataRoute.HandleFunc("/me/achievements", token.ValidateMiddleware(
		func(writer http.ResponseWriter, request *http.Request) {
			handleGetUserAchievements(writer, request, service)
		})).Methods("GET")
}

type Achievements struct {
	FirstCount int `json:"first_count"`
	SecondCount int `json:"second_count"`
	ThirdCount int `json:"third_count"`
	FinishedCount int `json:"finished_count"`
}

func handleGetUserAchievements(w http.ResponseWriter, request *http.Request, dataService data.IService) {
	user, err := token.GetUser(request, dataService)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if user == nil{
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var achievements Achievements
	achievements.FinishedCount, err = dataService.GetAthleteActivitiesByStateCount(user.Athlete.Id, models.ActivityStates.Finished)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	achievements.FirstCount, err = dataService.GetAthleteActivitiesPlaceCount(user.Athlete.Id, 1)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	achievements.SecondCount, err = dataService.GetAthleteActivitiesPlaceCount(user.Athlete.Id, 2)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	achievements.ThirdCount, err = dataService.GetAthleteActivitiesPlaceCount(user.Athlete.Id, 3)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	reply, _ := json.Marshal(achievements)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(reply)
	if err != nil {
		log.Error(err)
	}
	
}

func handleGetUser(w http.ResponseWriter, r *http.Request, dataService data.IService) {
	user, err := token.GetUser(r, dataService)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if user == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	reply, _ := json.Marshal(user.Athlete)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(reply)
	if err != nil {
		log.Error(err)
	}
}
