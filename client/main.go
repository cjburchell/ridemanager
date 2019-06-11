package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/cjburchell/tools-go/env"

	"github.com/cjburchell/go-uatu"
	"github.com/cjburchell/go-uatu/settings"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	err := settings.SetupLogger()
	if err != nil {
		log.Warn(err, "Unable to Connect to logger")
	}

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "ridemanager/dist/ridemanager/index.html")
	})

	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("ridemanager/dist/ridemanager"))))

	loggedRouter := handlers.LoggingHandler(log.Writer{Level: log.DEBUG}, r)

	port := env.GetInt("PORT", 8091)

	log.Printf("Starting Server at port %d", port)
	srv := &http.Server{
		Handler:      loggedRouter,
		Addr:         fmt.Sprintf(":%d", port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Error(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	err = srv.Shutdown(ctx)
	if err != nil {
		log.Error(err)
	}

	log.Print("shutting down")
	os.Exit(0)
}
