package settings_route

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cjburchell/go.strava"

	"github.com/cjburchell/tools-go/env"

	"github.com/cjburchell/go-uatu"
	"github.com/gorilla/mux"
)

// SetupDataRoute setup the route
func Setup(r *mux.Router) {
	dataRoute := r.PathPrefix("/api/v1/settings").Subrouter()
	dataRoute.HandleFunc("/{Setting}", handleGetSettings).Methods("GET")
}

func handleGetSettings(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	setting := vars["Setting"]

	result := ""
	switch setting {
	case "stravaClientId":
		result = fmt.Sprintf("%d", strava.ClientId)
	case "stravaRedirect":
		result = env.Get("STRAVA_REDIRECT_URI", "http://localhost:8091/api/v3/login")
	}

	if result == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	reply, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(reply)
	if err != nil {
		log.Error(err)
	}
}