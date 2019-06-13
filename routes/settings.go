package routes

import (
	"encoding/json"
	"net/http"

	"github.com/cjburchell/go-uatu"
	"github.com/gorilla/mux"
)

// SetupDataRoute setup the route
func SetupSettingsRoute(r *mux.Router) {
	dataRoute := r.PathPrefix("/api/v1/settings").Subrouter()
	dataRoute.HandleFunc("/{Setting}", handleGetSettings).Methods("GET")
}

func handleGetSettings(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	setting := vars["Setting"]

	reply, _ := json.Marshal(setting)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(reply)
	if err != nil {
		log.Error(err)
	}
}
