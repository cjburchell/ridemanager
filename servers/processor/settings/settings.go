package settings

import (
	"fmt"

	"github.com/cjburchell/settings-go"

	log "github.com/cjburchell/uatu-go"
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

func Get(logger log.ILog, settings settings.ISettings) (*Configuration, error) {
	config := &Configuration{
		MongoUrl:           settings.Get("MongoUrl", defaultMongoUrl),
		PollInterval:       settings.Get("PollInterval", defaultPollInterval),
		StravaClientId:     settings.GetInt("StravaClientId", 0),
		StravaClientSecret: settings.Get("StravaClientSecret", ""),
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
		warningMessage += fmt.Sprintf("\nMongoUrl set to default value (%s)", config.MongoUrl)
	}

	if warningMessage != "" {
		logger.Warn("Warning: " + warningMessage)
	}

	errorMessage := ""
	if config.StravaClientId == 0 {
		errorMessage += "\nStravaClientId Not set"
	}

	if config.StravaClientSecret == "" {
		errorMessage += "\nStravaClientSecret Not set"
	}

	if errorMessage != "" {
		logger.Error(nil, "ERRORS: "+errorMessage)
		return errors.New("Missing Env Settings")
	}

	return nil
}
