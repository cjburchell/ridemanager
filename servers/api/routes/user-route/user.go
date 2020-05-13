package user_route

import (
	"encoding/json"
	"net/http"

	"github.com/cjburchell/ridemanager/common/service/data/models"

	"github.com/cjburchell/uatu-go"

	"github.com/cjburchell/ridemanager/api/routes/token"

	"github.com/cjburchell/ridemanager/common/service/data"
	"github.com/gorilla/mux"
)

type handler struct {
	log log.ILog
}

func Setup(r *mux.Router, service data.IService, validator token.Validator, logger log.ILog) {
	dataRoute := r.PathPrefix("/api/v1/user").Subrouter()
	handle := handler{logger}
	dataRoute.HandleFunc("/me", validator.ValidateMiddleware(
		func(writer http.ResponseWriter, request *http.Request) {
			handle.getUser(writer, request, service)
		})).Methods("GET")

	dataRoute.HandleFunc("/me/achievements", validator.ValidateMiddleware(
		func(writer http.ResponseWriter, request *http.Request) {
			handle.getUserAchievements(writer, request, service)
		})).Methods("GET")
}

type Achievements struct {
	FirstCount    int `json:"first_count"`
	SecondCount   int `json:"second_count"`
	ThirdCount    int `json:"third_count"`
	FinishedCount int `json:"finished_count"`
}

func (h handler) getUserAchievements(w http.ResponseWriter, request *http.Request, dataService data.IService) {
	user, err := token.GetUser(request, dataService)
	if err != nil {
		h.log.Error(err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if user == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var achievements Achievements
	achievements.FinishedCount, err = dataService.GetAthleteActivitiesByStateCount(user.Athlete.Id, models.ActivityStates.Finished)
	if err != nil {
		h.log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	achievements.FirstCount, err = dataService.GetAthleteActivitiesPlaceCount(user.Athlete.Id, 1)
	if err != nil {
		h.log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	achievements.SecondCount, err = dataService.GetAthleteActivitiesPlaceCount(user.Athlete.Id, 2)
	if err != nil {
		h.log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	achievements.ThirdCount, err = dataService.GetAthleteActivitiesPlaceCount(user.Athlete.Id, 3)
	if err != nil {
		h.log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	reply, _ := json.Marshal(achievements)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(reply)
	if err != nil {
		h.log.Error(err)
	}

}

func (h handler) getUser(w http.ResponseWriter, r *http.Request, dataService data.IService) {
	user, err := token.GetUser(r, dataService)
	if err != nil {
		h.log.Error(err)
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
		h.log.Error(err)
	}
}
