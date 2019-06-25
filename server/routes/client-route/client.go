package client_route

import (
	"github.com/cjburchell/ridemanager/settings"
	"mime"
	"net/http"

	"github.com/cjburchell/go-uatu"
	"github.com/gorilla/mux"
)

func Setup(r *mux.Router) {


	err := mime.AddExtensionType(".js", "application/javascript; charset=utf-8")
	if err != nil {
		log.Error(err)
	}

	err = mime.AddExtensionType(".html", "text/html; charset=utf-8")
	if err != nil {
		log.Error(err)
	}

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, settings.ClientLocation+"/index.html")
	})

	r.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, settings.ClientLocation+"/index.html")
	})

	r.HandleFunc("/main", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, settings.ClientLocation+"/index.html")
	})

	r.HandleFunc("/main/{mode}", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, settings.ClientLocation+"/index.html")
	})

	r.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, settings.ClientLocation+"/index.html")
	})

	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(settings.ClientLocation))))
}
