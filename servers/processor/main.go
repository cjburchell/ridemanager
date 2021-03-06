package main

import (
	"os"
	"os/signal"

	"github.com/cjburchell/ridemanager/common/service/activity"
	"github.com/cjburchell/ridemanager/common/service/data"
	"github.com/cjburchell/ridemanager/common/service/data/models"
	"github.com/cjburchell/ridemanager/common/service/stravaService"
	appSettings "github.com/cjburchell/ridemanager/processor/settings"
	"github.com/cjburchell/settings-go"
	"github.com/cjburchell/tools-go/env"
	log "github.com/cjburchell/uatu-go"
	logSettings "github.com/cjburchell/uatu-go/settings"

	"github.com/robfig/cron"
)

func main() {
	set := settings.Get(env.Get("ConfigFile", "config.yaml"))
	logger := log.Create(logSettings.Get(set.GetSection("Logging")))

	config, err := appSettings.Get(logger, set)
	if err != nil {
		logger.Fatal(err, "Unable to verify settings")
	}

	dataService, err := data.NewService(data.GetSettings(set.GetSection("Data")))
	if err != nil {
		logger.Fatal(err, "Unable to Connect to mongo")
	}

	authenticator := stravaService.GetAuthenticator(config.StravaClientId, config.StravaClientSecret)



	cronTasks := startProcessor(config.PollInterval, dataService, authenticator, logger)
	defer cronTasks.Stop()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	logger.Print("shutting down")
	os.Exit(0)
}

func startProcessor(interval string, dataService data.IService, authenticator stravaService.Authenticator, logger log.ILog) *cron.Cron {
	cronTasks := cron.New()
	err := cronTasks.AddFunc(interval, func() {
		go updateAll(models.ActivityStates.Upcoming, dataService, authenticator, logger)
		go updateAll(models.ActivityStates.InProgress, dataService, authenticator, logger)
		go updateAll(models.ActivityStates.Finished, dataService, authenticator, logger)
	})

	if err != nil {
		logger.Error(err)
	}

	cronTasks.Start()

	return cronTasks
}

func updateAll(state models.ActivityState, service data.IService, authenticator stravaService.Authenticator, logger log.ILog) {
	activities, err := service.GetActivitiesByState(state)
	if err != nil {
		logger.Error(err, "Unable to get %s activities", state)
	}

	err = activity.UpdateAll(activities, service, true, authenticator)
	if err != nil {
		logger.Error(err, "Unable to update %s activities", state)
	}
}
