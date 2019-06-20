package activity_route

import (
	"net/http"

	log "github.com/cjburchell/go-uatu"
	"github.com/cjburchell/ridemanager/service/data/models"

	"github.com/cjburchell/ridemanager/routes/token"
	"github.com/cjburchell/ridemanager/service/data"
	"github.com/gorilla/mux"
)

func Setup(r *mux.Router, service data.IService) {

	dataRoute := r.PathPrefix("/api/v1/activity").Subrouter()
	dataRoute.HandleFunc("/{ActivityId}", token.ValidateMiddleware(
		func(writer http.ResponseWriter, request *http.Request) {
			handleGetActivity(writer, request, service)
		})).Methods("GET")

	dataRoute.HandleFunc("/", token.ValidateMiddleware(
		func(writer http.ResponseWriter, request *http.Request) {
			handleGetActivities(writer, request, service)
		})).Methods("GET")

	dataRoute.HandleFunc("/", token.ValidateMiddleware(
		func(writer http.ResponseWriter, request *http.Request) {
			handleCreateActivity(writer, request, service)
		})).Methods("POST")

	dataRoute.HandleFunc("/{ActivityId}", token.ValidateMiddleware(
		func(writer http.ResponseWriter, request *http.Request) {
			handleUpdateActivity(writer, request, service)
		})).Methods("PATCH")

	dataRoute.HandleFunc("/{ActivityId}", token.ValidateMiddleware(
		func(writer http.ResponseWriter, request *http.Request) {
			handleDeleteActivity(writer, request, service)
		})).Methods("DELETE")

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

}

func handleCreateActivity(writer http.ResponseWriter, request *http.Request, service data.IService) {

}

func handleGetActivities(writer http.ResponseWriter, request *http.Request, service data.IService) {

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

	writer.WriteHeader(http.StatusOK)

}
