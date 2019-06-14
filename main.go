package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/cjburchell/ridemanager/service/data"
	"github.com/cjburchell/ridemanager/settings"

	"github.com/cjburchell/go.strava"

	"github.com/cjburchell/ridemanager/routes"

	"github.com/robfig/cron"

	"github.com/cjburchell/go-uatu"
	logSettings "github.com/cjburchell/go-uatu/settings"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	err := logSettings.SetupLogger()
	if err != nil {
		log.Warn(err, "Unable to Connect to logger")
	}

	dataService, err := data.NewService(settings.MongoUrl)
	if err != nil {
		log.Warn(err, "Unable to Connect to mongo")
	}

	setupStrava()

	srv := startHttpServer(settings.Port, dataService)
	defer stopHttpServer(srv)

	cronTasks := startProcessor(settings.PollInterval)
	defer cronTasks.Stop()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	log.Print("shutting down")
	os.Exit(0)
}

func setupStrava() {
	strava.ClientId = settings.StravaClientId
	strava.ClientSecret = settings.StravaClientSecret
}

func startProcessor(interval string) *cron.Cron {
	cronTasks := cron.New()
	err := cronTasks.AddFunc(interval, func() {

	})

	if err != nil {
		log.Error(err)
	}

	cronTasks.Start()

	return cronTasks
}

func startHttpServer(port int, service data.IService) *http.Server {
	r := mux.NewRouter()
	routes.SetupLoginRoute(r, service)
	routes.SetupStatusRoute(r)
	routes.SetupSettingsRoute(r)
	routes.SetupClientRoute(r)

	loggedRouter := handlers.LoggingHandler(log.Writer{Level: log.DEBUG}, r)

	log.Printf("Starting Server at port %d", port)
	srv := &http.Server{
		Handler:      loggedRouter,
		Addr:         fmt.Sprintf(":%d", port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error(err)
		}
	}()

	return srv
}

func stopHttpServer(srv *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	err := srv.Shutdown(ctx)
	if err != nil {
		log.Error(err)
	}
}
