package settings

import (
	"fmt"
	"github.com/pkg/errors"

	log "github.com/cjburchell/go-uatu"
	"github.com/cjburchell/tools-go/env"
)

const defaultPort = 8091
const defaultClientLocation = "client/dist/ridemanager-client"
const defaultApiAddress = "http://localhost:8091/api/v1"

type Configuration struct {
	Port               int
	ClientLocation 	   string
	ApiAddress 	       string
	StravaClientId     int
	MapboxToken        string
}

func Get(logger log.ILog) (*Configuration, error) {
	config := &Configuration{
		Port:              env.GetInt("PORT", defaultPort),
		ClientLocation:    env.Get("CLIENT_LOCATION", defaultClientLocation),
		MapboxToken:       env.Get("MAPBOX_ACCESS_TOKEN", ""),
		ApiAddress:        env.Get("API_LOCATION", defaultApiAddress),
	}

	err := config.verify(logger)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func (config Configuration) verify(logger log.ILog) error {

	warningMessage := ""
	if config.ClientLocation == defaultClientLocation {
		warningMessage += fmt.Sprintf("\nCLIENT_LOCATION set to default value (%s)", config.ClientLocation)
	}

	if config.ApiAddress == defaultApiAddress {
		warningMessage += fmt.Sprintf("\nAPI_LOCATION set to default value (%s)", config.ApiAddress)
	}

	if warningMessage != "" {
		logger.Warn("Warning: " + warningMessage)
	}

	errorMessage := ""
	if config.MapboxToken == "" {
		errorMessage += "\nMAPBOX_ACCESS_TOKEN Not set"
	}

	if errorMessage != "" {
		logger.Error(nil, "ERRORS: "+errorMessage)
		return errors.New("Missing Env Settings")
	}

	return nil
}
