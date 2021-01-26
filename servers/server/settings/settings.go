package settings

import (
	"fmt"

	"github.com/cjburchell/settings-go"

	log "github.com/cjburchell/uatu-go"
	"github.com/pkg/errors"
)

const defaultPort = 8091
const defaultJwtSecret = "test"
const defaultClientLocation = "client/dist/ridemanager-client"
const defaultStravaRedirect = "http://localhost:8091/api/v1/login/validate"

// Configuration of the server
type Configuration struct {
	Port                int
	StravaClientID      int
	StravaClientSecret  string
	JwtSecret           string
	ClientLocation      string
	MapboxToken         string
	StravaLoginRedirect string
}

// Get the server settings
func Get(logger log.ILog, settings settings.ISettings) (*Configuration, error) {
	config := &Configuration{
		Port:                settings.GetInt("Port", defaultPort),
		StravaClientID:      settings.GetInt("StravaClientId", 0),
		StravaClientSecret:  settings.Get("StravaClientSecret", ""),
		JwtSecret:           settings.Get("JwtSecret", defaultJwtSecret),
		ClientLocation:      settings.Get("ClientLocation", defaultClientLocation),
		MapboxToken:         settings.Get("MapboxToken", ""),
		StravaLoginRedirect: settings.Get("StravaLoginRedirectURL", defaultStravaRedirect),
	}

	err := config.verify(logger)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func (config Configuration) verify(logger log.ILog) error {

	warningMessage := ""
	if config.JwtSecret == defaultJwtSecret {
		warningMessage += fmt.Sprintf("\nJwtSecret set to default value (%s)", config.JwtSecret)
	}

	if config.ClientLocation == defaultClientLocation {
		warningMessage += fmt.Sprintf("\nClientLocation set to default value (%s)", config.ClientLocation)
	}

	if config.StravaLoginRedirect == defaultStravaRedirect {
		warningMessage += fmt.Sprintf("\nStravaLoginRedirectURL set to default value (%s)", config.StravaLoginRedirect)
	}

	if warningMessage != "" {
		logger.Warn("Warning: " + warningMessage)
	}

	errorMessage := ""

	if config.StravaClientID == 0 {
		errorMessage += "\nSTRAVA_CLIENT_ID Not set"
	}

	if config.StravaClientSecret == "" {
		errorMessage += "\nSTRAVA_CLIENT_SECRET Not set"
	}

	if config.MapboxToken == "" {
		errorMessage += "\nMAPBOX_ACCESS_TOKEN Not set"
	}

	if errorMessage != "" {
		logger.Error(nil, "ERRORS: "+errorMessage)
		return errors.New("Missing Env Settings")
	}

	return nil
}
