package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/cjburchell/ridemanager/routes"

	"github.com/robfig/cron"

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

	srv := startHttpServer(env.GetInt("PORT", 8091))
	defer stopHttpServer(srv)

	cronTasks := startProcessor(env.Get("POLL_INTERVAL", "@hourly"))
	defer cronTasks.Stop()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	log.Print("shutting down")
	os.Exit(0)
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

func startHttpServer(port int) *http.Server {
	r := mux.NewRouter()
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
