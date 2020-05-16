package settings

import (
	"fmt"

	"github.com/cjburchell/settings-go"

	log "github.com/cjburchell/uatu-go"
	"github.com/pkg/errors"
)

const defaultMongoUrl = "localhost"
const defaultPort = 8091
const defaultJwtSecret = "test"
const defaultClientLocation = "client/dist/ridemanager-client"
const defaultStravaRedirect = "http://localhost:8091/api/v1/login/validate"

type Configuration struct {
	MongoUrl            string
	Port                int
	StravaClientId      int
	StravaClientSecret  string
	JwtSecret           string
	ClientLocation      string
	MapboxToken         string
	StravaLoginRedirect string
}

func Get(logger log.ILog, settings settings.ISettings) (*Configuration, error) {
	config := &Configuration{
		MongoUrl:            settings.Get("MONGO_URL", defaultMongoUrl),
		Port:                settings.GetInt("PORT", defaultPort),
		StravaClientId:      settings.GetInt("STRAVA_CLIENT_ID", 0),
		StravaClientSecret:  settings.Get("STRAVA_CLIENT_SECRET", ""),
		JwtSecret:           settings.Get("JWT_SECRET", defaultJwtSecret),
		ClientLocation:      settings.Get("CLIENT_LOCATION", defaultClientLocation),
		MapboxToken:         settings.Get("MAPBOX_ACCESS_TOKEN", ""),
		StravaLoginRedirect: settings.Get("STRAVA_LOGIN_REDIRECT_URL", defaultStravaRedirect),
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

	if config.JwtSecret == defaultJwtSecret {
		warningMessage += fmt.Sprintf("\nJWT_SECRET set to default value (%s)", config.JwtSecret)
	}

	if config.ClientLocation == defaultClientLocation {
		warningMessage += fmt.Sprintf("\nCLIENT_LOCATION set to default value (%s)", config.ClientLocation)
	}

	if config.StravaLoginRedirect == defaultStravaRedirect {
		warningMessage += fmt.Sprintf("\nSTRAVA_LOGIN_REDIRECT_URL set to default value (%s)", config.StravaLoginRedirect)
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

	if config.MapboxToken == "" {
		errorMessage += "\nMAPBOX_ACCESS_TOKEN Not set"
	}

	if errorMessage != "" {
		logger.Error(nil, "ERRORS: "+errorMessage)
		return errors.New("Missing Env Settings")
	}

	return nil
}
