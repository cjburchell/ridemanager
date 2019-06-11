package main

import (
	"os"
	"os/signal"

	"github.com/cjburchell/go-uatu"
	"github.com/cjburchell/go-uatu/settings"
	"github.com/cjburchell/tools-go/env"
	"github.com/robfig/cron"
)

func main() {
	err := settings.SetupLogger()
	if err != nil {
		log.Warn(err, "Unable to Connect to logger")
	}

	interval := env.Get("POLL_INTERVAL", "@hourly")

	cronTasks := cron.New()
	err = cronTasks.AddFunc(interval, func() {

	})

	if err != nil {
		log.Error(err)
	}

	cronTasks.Start()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	cronTasks.Stop()

	log.Print("shutting down")
	os.Exit(0)
}
