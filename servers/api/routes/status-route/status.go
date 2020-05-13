package status_route

import (
	"net/http"

	"github.com/cjburchell/uatu-go"
	"github.com/gorilla/mux"
)

type handler struct {
	log log.ILog
}

func Setup(r *mux.Router, logger log.ILog) {
	handle := handler{logger}
	r.HandleFunc("/@status", handle.getStatus).Methods("GET")
}

func (h handler) getStatus(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write([]byte("Ok"))
	if err != nil {
		h.log.Error(err)
	}
}
