package settings

import (
	"fmt"

	log "github.com/cjburchell/go-uatu"
	"github.com/cjburchell/tools-go/env"
	"github.com/pkg/errors"
)

const defaultMongoUrl = "localhost"
const defaultPort = 8091
const defaultStravaRedirectURI = "http://localhost:8091/api/v1/login"
const defaultJwtSecret = "test"

type Configuration struct {
	MongoUrl           string
	Port               int
	StravaClientId     int
	StravaClientSecret string
	StravaRedirectURI  string
	JwtSecret          string
	MapboxToken        string
}

func Get(logger log.ILog) (*Configuration, error) {
	config := &Configuration{
		MongoUrl:           env.Get("MONGO_URL", defaultMongoUrl),
		Port:               env.GetInt("PORT", defaultPort),
		StravaClientId:     env.GetInt("STRAVA_CLIENT_ID", 0),
		StravaClientSecret: env.Get("STRAVA_CLIENT_SECRET", ""),
		StravaRedirectURI:  env.Get("STRAVA_REDIRECT_URI", defaultStravaRedirectURI),
		JwtSecret:          env.Get("JWT_SECRET", defaultJwtSecret),
		MapboxToken:        env.Get("MAPBOX_ACCESS_TOKEN", ""),
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

	if config.StravaRedirectURI == defaultStravaRedirectURI {
		warningMessage += fmt.Sprintf("\nSTRAVA_REDIRECT_URI set to default value (%s)", config.StravaRedirectURI)
	}

	if config.JwtSecret == defaultJwtSecret {
		warningMessage += fmt.Sprintf("\nJWT_SECRET set to default value (%s)", config.JwtSecret)
	}

	if warningMessage != "" {
		logger.Warn("Warning: " + warningMessage)
	}

	errorMessage := ""
	if config.MapboxToken == "" {
		errorMessage += "\nMAPBOX_ACCESS_TOKEN Not set"
	}

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
