package settings

import (
	"fmt"

	log "github.com/cjburchell/go-uatu"
	"github.com/cjburchell/tools-go/env"
	"github.com/pkg/errors"
)

const defaultMongoUrl = "localhost"
const defaultPollInterval = "@hourly"

type Configuration struct {
	MongoUrl           string
	PollInterval       string
	StravaClientId     int
	StravaClientSecret string
}

func Get(logger log.ILog) (*Configuration, error) {
	config := &Configuration{
		MongoUrl:           env.Get("MONGO_URL", defaultMongoUrl),
		PollInterval:       env.Get("POLL_INTERVAL", defaultPollInterval),
		StravaClientId:     env.GetInt("STRAVA_CLIENT_ID", 0),
		StravaClientSecret: env.Get("STRAVA_CLIENT_SECRET", ""),
	}

	err := config.verify(logger)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func (config Configuration) verify(logger log.ILog) error {

	warningMessage := ""
	if config.MongoUrl == defaultMongoUrl {
		warningMessage += fmt.Sprintf("\nMONGO_URL set to default value (%s)", config.MongoUrl)
	}

	if warningMessage != "" {
		logger.Warn("Warning: " + warningMessage)
	}

	errorMessage := ""
	if config.StravaClientId == 0 {
		errorMessage += "\nSTRAVA_CLIENT_ID Not set"
	}

	if config.StravaClientSecret == "" {
		errorMessage += "\nSTRAVA_CLIENT_SECRET Not set"
	}

	if errorMessage != "" {
		logger.Error(nil, "ERRORS: "+errorMessage)
		return errors.New("Missing Env Settings")
	}

	return nil
}
