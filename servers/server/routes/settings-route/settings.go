package settings_route

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cjburchell/ridemanager/server/settings"

	log "github.com/cjburchell/uatu-go"
	"github.com/gorilla/mux"
)

type handler struct {
	settings.Configuration
	log log.ILog
}

// SetupDataRoute setup the route
func Setup(r *mux.Router, configuration settings.Configuration, logger log.ILog) {
	dataRoute := r.PathPrefix("/client/settings").Subrouter()
	handle := handler{configuration, logger}
	dataRoute.HandleFunc("/{Setting}", handle.getSettings).Methods("GET")
}

func (h handler) getSettings(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	setting := vars["Setting"]

	result := ""
	switch setting {
	case "stravaClientId":
		result = fmt.Sprintf("%d", h.StravaClientID)
	case "stravaRedirect":
		result = h.StravaLoginRedirect
	case "mapboxAccessToken":
		result = h.MapboxToken
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
		h.log.Error(err)
	}
}
