package routes

import (
	"mime"
	"net/http"

	"github.com/cjburchell/tools-go/env"

	"github.com/cjburchell/go-uatu"
	"github.com/gorilla/mux"
)

func SetupClientRoute(r *mux.Router) {
	clientLocation := env.Get("CLIENT_LOCATION", "ridemanager-client/dist/ridemanager-client")

	err := mime.AddExtensionType(".js", "application/javascript; charset=utf-8")
	if err != nil {
		log.Error(err)
	}

	err = mime.AddExtensionType(".html", "text/html; charset=utf-8")
	if err != nil {
		log.Error(err)
	}

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, clientLocation+"/index.html")
	})

	r.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, clientLocation+"/index.html")
	})

	r.HandleFunc("/main", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, clientLocation+"/index.html")
	})

	r.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, clientLocation+"/index.html")
	})

	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(clientLocation))))
}
