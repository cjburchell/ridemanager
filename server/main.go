package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/cjburchell/ridemanager/service/stravaService"

	"github.com/cjburchell/ridemanager/routes/token"

	activityRoute "github.com/cjburchell/ridemanager/routes/activity-route"
	clientRoute "github.com/cjburchell/ridemanager/routes/client-route"
	loginRoute "github.com/cjburchell/ridemanager/routes/login-route"
	settingsRoute "github.com/cjburchell/ridemanager/routes/settings-route"
	statusRoute "github.com/cjburchell/ridemanager/routes/status-route"
	stravaRoute "github.com/cjburchell/ridemanager/routes/strava-route"
	userRoute "github.com/cjburchell/ridemanager/routes/user-route"

	"github.com/cjburchell/ridemanager/service/data"
	"github.com/cjburchell/ridemanager/settings"

	"github.com/robfig/cron"

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

	dataService, err := data.NewService(config.MongoUrl)
	if err != nil {
		logger.Fatal(err, "Unable to Connect to mongo")
	}

	srv := startHTTPServer(*config, dataService, logger)
	defer stopHTTPServer(srv, logger)

	cronTasks := startProcessor(config.PollInterval, logger)
	defer cronTasks.Stop()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	logger.Print("shutting down")
	os.Exit(0)
}

func startProcessor(interval string, logger log.ILog) *cron.Cron {
	cronTasks := cron.New()
	err := cronTasks.AddFunc(interval, func() {

	})

	if err != nil {
		logger.Error(err)
	}

	cronTasks.Start()

	return cronTasks
}

func startHTTPServer(config settings.Configuration, service data.IService, logger log.ILog) *http.Server {
	r := mux.NewRouter()

	tokenValidator := token.GetValidator(config.JwtSecret)
	tokenBuilder := token.GetBuilder(config.JwtSecret)
	authenticator := stravaService.GetAuthenticator(config.StravaClientId, config.StravaClientSecret)
	loginRoute.Setup(r, service, tokenValidator, tokenBuilder, authenticator, logger)
	userRoute.Setup(r, service, tokenValidator, logger)
	activityRoute.Setup(r, service, tokenValidator, authenticator, logger)
	stravaRoute.Setup(r, service, tokenValidator, authenticator, logger)
	statusRoute.Setup(r, logger)
	settingsRoute.Setup(r, config, logger)
	clientRoute.Setup(r, config.ClientLocation, logger)

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
