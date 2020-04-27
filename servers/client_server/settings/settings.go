package settings

import (
	"fmt"

	log "github.com/cjburchell/go-uatu"
	"github.com/cjburchell/tools-go/env"
)

const defaultPort = 8091
const defaultClientLocation = "client/dist/ridemanager-client"

type Configuration struct {
	Port           int
	ClientLocation string
}

func Get(logger log.ILog) (*Configuration, error) {
	config := &Configuration{
		Port:           env.GetInt("PORT", defaultPort),
		ClientLocation: env.Get("CLIENT_LOCATION", defaultClientLocation),
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

	if warningMessage != "" {
		logger.Warn("Warning: " + warningMessage)
	}

	return nil
}
