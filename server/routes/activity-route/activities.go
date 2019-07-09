package activity_route

import (
	"encoding/json"
	"github.com/cjburchell/go-uatu"
	"github.com/cjburchell/ridemanager/service/data/models"
	"net/http"

	"github.com/cjburchell/ridemanager/routes/token"
	"github.com/cjburchell/ridemanager/service/data"
	"github.com/gorilla/mux"
)

func Setup(r *mux.Router, service data.IService) {

	dataRoute := r.PathPrefix("/api/v1/activity").Subrouter()

	dataRoute.HandleFunc("/my", token.ValidateMiddleware(
		func(writer http.ResponseWriter, request *http.Request) {
			handleGetMyActivities(writer, request, service)
		})).Methods("GET")

	dataRoute.HandleFunc("/public", token.ValidateMiddleware(
		func(writer http.ResponseWriter, request *http.Request) {
			handleGetPublicActivities(writer, request, service)
		})).Methods("GET")

	dataRoute.HandleFunc("/joined", token.ValidateMiddleware(
		func(writer http.ResponseWriter, request *http.Request) {
			handleGetJoinedActivities(writer, request, service)
		})).Methods("GET")

	dataRoute.HandleFunc("/{ActivityId}", token.ValidateMiddleware(
		func(writer http.ResponseWriter, request *http.Request) {
			handleGetActivity(writer, request, service)
		})).Methods("GET")

	dataRoute.HandleFunc("/", token.ValidateMiddleware(
		func(writer http.ResponseWriter, request *http.Request) {
			handleCreateActivity(writer, request, service)
		})).Methods("POST")

	dataRoute.HandleFunc("/{ActivityId}", token.ValidateMiddleware(validateWritableAccessMiddleware(
		func(writer http.ResponseWriter, request *http.Request) {
			handleUpdateActivity(writer, request, service)
		}, service))).Methods("PATCH")

	dataRoute.HandleFunc("/{ActivityId}", token.ValidateMiddleware(validateWritableAccessMiddleware(
		func(writer http.ResponseWriter, request *http.Request) {
			handleDeleteActivity(writer, request, service)
		}, service))).Methods("DELETE")

}

func validateWritableAccessMiddleware(next http.HandlerFunc, service data.IService) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		 user, err := token.GetUser(req, service)
		 if err != nil{
			 w.WriteHeader(http.StatusUnauthorized)
			 return
		 }

		if user == nil{
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// check user role
		if user.Role == models.AdminRole{
			next(w, req)
			return
		}

		 // check activity owner
		vars := mux.Vars(req)
		activityId := models.ActivityId(vars["ActivityId"])

		activity, err := service.GetActivity(activityId)
		if err != nil {
			log.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if activity.OwnerId != user.Id{
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next(w, req)
	})
}

func handleDeleteActivity(writer http.ResponseWriter, request *http.Request, service data.IService) {
	vars := mux.Vars(request)
	activityId := models.ActivityId(vars["ActivityId"])

	err := service.DeleteActivity(activityId)
	if err != nil {
		log.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
}

func handleUpdateActivity(writer http.ResponseWriter, request *http.Request, service data.IService) {
	decoder := json.NewDecoder(request.Body)
	var activity models.Activity
	err := decoder.Decode(&activity)
	if err != nil {
		log.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = service.UpdateActivity(activity)
	if err != nil {
		log.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
}

func handleCreateActivity(writer http.ResponseWriter, request *http.Request, service data.IService) {
	decoder := json.NewDecoder(request.Body)
	var activity models.Activity
	err := decoder.Decode(&activity)
	if err != nil {
		log.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = service.AddActivity(&activity)
	if err != nil {
		log.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
}

func handleGetMyActivities(writer http.ResponseWriter, req *http.Request, service data.IService) {

	user, err := token.GetUser(req, service)
	if err != nil{
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	activities, err := service.GetOwnedActivities(user.Id)
	if err != nil {
		log.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	reply, _ := json.Marshal(activities)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	_, err = writer.Write(reply)
	if err != nil {
		log.Error(err)
	}
}

func handleGetJoinedActivities(writer http.ResponseWriter, req *http.Request, service data.IService) {

	user, err := token.GetUser(req, service)
	if err != nil{
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	activities, err := service.GetAthleteActivities(user.Id)
	if err != nil {
		log.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	reply, _ := json.Marshal(activities)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	_, err = writer.Write(reply)
	if err != nil {
		log.Error(err)
	}
}

func handleGetPublicActivities(writer http.ResponseWriter, _ *http.Request, service data.IService) {
	activities, err := service.GetActivitiesByPrivacy(models.PublicActivity)
	if err != nil {
		log.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	reply, _ := json.Marshal(activities)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	_, err = writer.Write(reply)
	if err != nil {
		log.Error(err)
	}
}

func handleGetActivity(writer http.ResponseWriter, request *http.Request, service data.IService) {

	vars := mux.Vars(request)
	activityId := models.ActivityId(vars["ActivityId"])

	activity, err := service.GetActivity(activityId)
	if err != nil {
		log.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	reply, _ := json.Marshal(activity)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	_, err = writer.Write(reply)
	if err != nil {
		log.Error(err)
	}

}
