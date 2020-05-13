package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	clientRoute "github.com/cjburchell/ridemanager/api/routes/client-route"
	settingsRoute "github.com/cjburchell/ridemanager/api/routes/settings-route"
	"github.com/cjburchell/settings-go"
	"github.com/cjburchell/tools-go/env"

	"github.com/cjburchell/ridemanager/common/service/stravaService"

	"github.com/cjburchell/ridemanager/api/routes/token"

	activityRoute "github.com/cjburchell/ridemanager/api/routes/activity-route"
	loginRoute "github.com/cjburchell/ridemanager/api/routes/login-route"
	statusRoute "github.com/cjburchell/ridemanager/api/routes/status-route"
	stravaRoute "github.com/cjburchell/ridemanager/api/routes/strava-route"
	userRoute "github.com/cjburchell/ridemanager/api/routes/user-route"
	serverSettings "github.com/cjburchell/ridemanager/api/settings"
	"github.com/cjburchell/ridemanager/common/service/data"

	log "github.com/cjburchell/uatu-go"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	set := settings.Get(env.Get("ConfigFile", ""))
	logger := log.Create(set)

	config, err := serverSettings.Get(logger, set)
	if err != nil {
		logger.Fatal(err, "Unable to verify settings")
	}

	dataService, err := data.NewService(config.MongoUrl)
	if err != nil {
		logger.Fatal(err, "Unable to Connect to mongo")
	}

	srv := startHTTPServer(*config, dataService, logger)
	defer stopHTTPServer(srv, logger)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	logger.Print("shutting down")
	os.Exit(0)
}

func startHTTPServer(config serverSettings.Configuration, service data.IService, logger log.ILog) *http.Server {
	r := mux.NewRouter()

	tokenValidator := token.GetValidator(config.JwtSecret)
	tokenBuilder := token.GetBuilder(config.JwtSecret)
	authenticator := stravaService.GetAuthenticator(config.StravaClientId, config.StravaClientSecret)
	loginRoute.Setup(r, service, tokenValidator, tokenBuilder, authenticator, logger)
	userRoute.Setup(r, service, tokenValidator, logger)
	activityRoute.Setup(r, service, tokenValidator, authenticator, logger)
	stravaRoute.Setup(r, service, tokenValidator, authenticator, logger)
	statusRoute.Setup(r, logger)
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
