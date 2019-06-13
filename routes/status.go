package routes

import (
	"net/http"

	"github.com/cjburchell/go-uatu"
	"github.com/gorilla/mux"
)

func SetupStatusRoute(r *mux.Router) {
	r.HandleFunc("/@status", handleGetStatus).Methods("GET")
}

func handleGetStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write([]byte("Ok"))
	if err != nil {
		log.Error(err)
	}
}
