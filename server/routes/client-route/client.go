package clientroute

import (
	"mime"
	"net/http"

	"github.com/cjburchell/go-uatu"
	"github.com/gorilla/mux"
)

func Setup(r *mux.Router, clientLocation string, logger log.ILog) {

	err := mime.AddExtensionType(".js", "application/javascript; charset=utf-8")
	if err != nil {
		logger.Error(err)
	}

	err = mime.AddExtensionType(".html", "text/html; charset=utf-8")
	if err != nil {
		logger.Error(err)
	}

	handleClient := func (w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, clientLocation+"/index.html")
	}

	r.HandleFunc("/", handleClient)
	r.HandleFunc("/login", handleClient)
	r.HandleFunc("/main", handleClient)
	r.HandleFunc("/main/{mode}", handleClient)
	r.HandleFunc("/activity/{activityId}", handleClient)
	r.HandleFunc("/edit/{activityId}", handleClient)
	r.HandleFunc("/token", handleClient)
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(clientLocation))))
}
