package status_route

import (
	"net/http"

	"github.com/cjburchell/go-uatu"
	"github.com/gorilla/mux"
)

type handler struct {
	log log.ILog
}

func Setup(r *mux.Router, logger log.ILog) {
	handle := handler{logger}
	r.HandleFunc("/@status", handle.getStatus).Methods("GET")
}

func (h handler) getStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write([]byte("Ok"))
	if err != nil {
		h.log.Error(err)
	}
}