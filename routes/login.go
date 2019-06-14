package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cjburchell/ridemanager/service/data/models"

	"github.com/cjburchell/ridemanager/service/data"

	"github.com/cjburchell/go-uatu"

	"github.com/cjburchell/go.strava"

	"github.com/gorilla/mux"
)

var authenticator = &strava.OAuthAuthenticator{}
var dataService data.IService

// SetupDataRoute setup the route
func SetupLoginRoute(r *mux.Router, service data.IService) {
	dataService = service
	dataRoute := r.PathPrefix("/api/v1/login").Subrouter()
	dataRoute.HandleFunc("/validate", authenticator.HandlerFunc(oAuthSuccess, oAuthFailure)).Methods("GET").Queries("code", "{code}", "scope", "{scope}")

	dataRoute.HandleFunc("/status", validateTokenMiddleware(getLoginStatus)).Methods("GET")
}

func oAuthSuccess(auth *strava.AuthorizationResponse, w http.ResponseWriter, r *http.Request) {
	log.Debugf("Access Token: %s", auth.AccessToken)

	user, err := dataService.GetStravaUser(auth.Athlete.Id)
	if err != nil {
		log.Error(err)
	}

	name := fmt.Sprintf("%s %s", auth.Athlete.FirstName, auth.Athlete.LastName)
	if user == nil {
		user = models.NewUser(auth.Athlete.Id)
		user.StravaToken = auth.AccessToken
		user.Name = name
		user.ProfileMediumImage = auth.Athlete.ProfileMedium
		user.ProfileImage = auth.Athlete.Profile
		user.Gender = auth.Athlete.Gender
		err = dataService.AddUser(user)
		if err != nil {
			log.Error(err)
		}
	} else {
		user.StravaToken = auth.AccessToken
		user.Name = name
		user.ProfileMediumImage = auth.Athlete.ProfileMedium
		user.ProfileImage = auth.Athlete.Profile
		user.Gender = auth.Athlete.Gender
		err = dataService.UpdateUser(*user)
		if err != nil {
			log.Error(err)
		}
	}

	tokenString, err := buildToken(user.Id)
	if err != nil {
		log.Error(err)
	}

	http.Redirect(w, r, "/token?token="+tokenString, http.StatusFound)
}

func oAuthFailure(err error, w http.ResponseWriter, r *http.Request) {
	log.Debugf("Authorization Failure:\n")

	// some standard error checking
	if err == strava.OAuthAuthorizationDeniedErr {
		log.Debugf("The user clicked the 'Do not Authorize' button on the previous page.\n")
		log.Debugf("This is the main error your application should handle.")
	} else if err == strava.OAuthInvalidCredentialsErr {
		log.Debugf("You provided an incorrect client_id or client_secret.\nDid you remember to set them at the begininng of this file?")
	} else if err == strava.OAuthInvalidCodeErr {
		log.Debugf("The temporary token was not recognized, this shouldn't happen normally")
	} else if err == strava.OAuthServerErr {
		log.Debugf("There was some sort of server error, try again to see if the problem continues")
	} else {
		log.Error(err)
	}

	http.Redirect(w, r, "/login", http.StatusFound)
}

func getLoginStatus(w http.ResponseWriter, r *http.Request) {
	user, err := getUser(r, dataService)
	if err != nil {
		log.Error(err)
	}

	log.Debugf("name %s,  userId: %s", user.Name, user.Id)

	reply, _ := json.Marshal(true)
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(reply)
	if err != nil {
		log.Error(err)
	}
}
