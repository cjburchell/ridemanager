package user_route

import (
	"encoding/json"
	"net/http"

	"github.com/cjburchell/go-uatu"

	"github.com/cjburchell/ridemanager/routes/contract"

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
}

func handleGetUser(w http.ResponseWriter, r *http.Request, dataService data.IService) {
	user, err := token.GetUser(r, dataService)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var contractUser = contract.NewUser(*user)

	reply, _ := json.Marshal(contractUser)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(reply)
	if err != nil {
		log.Error(err)
	}
}
