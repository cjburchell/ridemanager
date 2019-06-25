package settings

import "github.com/cjburchell/tools-go/env"

var MongoUrl = env.Get("MONGO_URL", "localhost")
var Port = env.GetInt("PORT", 8091)
var PollInterval = env.Get("POLL_INTERVAL", "@hourly")

var StravaClientId = env.GetInt("STRAVA_CLIENT_ID", 0)
var StravaClientSecret = env.Get("STRAVA_CLIENT_SECRET", "")

var StravaRedirectURI = env.Get("STRAVA_REDIRECT_URI", "http://localhost:8091/api/v1/login")

var JwtSecret = env.Get("JWT_SECRET", "test")

var ClientLocation = env.Get("CLIENT_LOCATION", "client/dist/ridemanager-client")
