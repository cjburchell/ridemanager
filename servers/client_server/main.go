package main

import (
	"context"
	"fmt"
	settingsRoute "github.com/cjburchell/ridemanager/client_server/routes/settings-route"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/cjburchell/ridemanager/client_server/settings"

	clientRoute "github.com/cjburchell/ridemanager/client_server/routes/client-route"

	"github.com/cjburchell/go-uatu"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	logger := log.Create()

	config, err := settings.Get(logger)
	if err != nil {
		logger.Fatal(err, "Unable to verify settings")
	}

	srv := startHTTPServer(*config, logger)
	defer stopHTTPServer(srv, logger)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	logger.Print("shutting down")
	os.Exit(0)
}

func startHTTPServer(config settings.Configuration, logger log.ILog) *http.Server {
	r := mux.NewRouter()

	clientRoute.Setup(r, config.ClientLocation, logger)
	settingsRoute.Setup(r, config, logger)

	loggedRouter := handlers.LoggingHandler(log.Writer{Level: log.DEBUG}, r)

	logger.Printf("Starting Server at port %d", config.Port)
	srv := &http.Server{
		Handler:      loggedRouter,
		Addr:         fmt.Sprintf(":%d", config.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error(err)
		}
	}()

	return srv
}

func stopHTTPServer(srv *http.Server, logger log.ILog) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	err := srv.Shutdown(ctx)
	if err != nil {
		logger.Error(err)
	}
}
