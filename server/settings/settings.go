package settings

import (
	"fmt"
	log "github.com/cjburchell/go-uatu"
	"github.com/cjburchell/tools-go/env"
	"github.com/pkg/errors"
)

const defaultMongoUrl ="localhost"
var MongoUrl = env.Get("MONGO_URL", defaultMongoUrl)

const defaultPort  = 8091
var Port = env.GetInt("PORT", defaultPort)

const defaultPollInterval  = "@hourly"
var PollInterval = env.Get("POLL_INTERVAL", defaultPollInterval)

var StravaClientId = env.GetInt("STRAVA_CLIENT_ID", 0)
var StravaClientSecret = env.Get("STRAVA_CLIENT_SECRET", "")

const defaultStravaRedirectURI = "http://localhost:8091/api/v1/login"
var StravaRedirectURI = env.Get("STRAVA_REDIRECT_URI", defaultStravaRedirectURI)

const defaultJwtSecret = "test"
var JwtSecret = env.Get("JWT_SECRET", defaultJwtSecret)

const defaultClientLocation  = "client/dist/ridemanager-client"
var ClientLocation = env.Get("CLIENT_LOCATION", defaultClientLocation)

var MapboxToken = env.Get("MAPBOX_ACCESS_TOKEN", "")

func Verify() error {

	warningMessage := ""
	if MongoUrl == defaultMongoUrl {
		warningMessage += fmt.Sprintf("\nMONGO_URL set to default value (%s)", MongoUrl)
	}

	if StravaRedirectURI == defaultStravaRedirectURI {
		warningMessage += fmt.Sprintf("\nSTRAVA_REDIRECT_URI set to default value (%s)", StravaRedirectURI)
	}

	if ClientLocation == defaultClientLocation {
		warningMessage += fmt.Sprintf("\nCLIENT_LOCATION set to default value (%s)", ClientLocation)
	}

	if JwtSecret == defaultJwtSecret {
		warningMessage += fmt.Sprintf("\nJWT_SECRET set to default value (%s)", JwtSecret)
	}

	if warningMessage != "" {
		log.Warn("Warning: " + warningMessage)
	}

	errorMessage := ""
	if MapboxToken == "" {
		errorMessage += "\nMAPBOX_ACCESS_TOKEN Not set"
	}

	if StravaClientId == 0 {
		errorMessage += "\nSTRAVA_CLIENT_ID Not set"
	}

	if StravaClientSecret == "" {
		errorMessage += "\nSTRAVA_CLIENT_SECRET Not set"
	}

	if errorMessage != "" {
		log.Error(nil, "ERRORS: "+errorMessage)
		return errors.New("Missing Env Settings")
	}

	return nil
}