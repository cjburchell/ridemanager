package login_route

import (
	"encoding/json"
	"net/http"

	"github.com/cjburchell/ridemanager/routes/token"

	"github.com/cjburchell/ridemanager/service/data/models"

	"github.com/cjburchell/ridemanager/service/data"

	"github.com/cjburchell/go-uatu"

	"github.com/cjburchell/go.strava"

	"github.com/gorilla/mux"
)

var authenticator = &strava.OAuthAuthenticator{}

// SetupDataRoute setup the route
func Setup(r *mux.Router, service data.IService) {
	dataRoute := r.PathPrefix("/api/v1/login").Subrouter()
	dataRoute.HandleFunc("/validate", authenticator.HandlerFunc(func(auth *strava.AuthorizationResponse, w http.ResponseWriter, r *http.Request) {
		oAuthSuccess(auth, w, r, service)
	}, oAuthFailure)).Methods("GET").Queries("code", "{code}", "scope", "{scope}")

	dataRoute.HandleFunc("/status", token.ValidateMiddleware(func(writer http.ResponseWriter, request *http.Request) {
		getLoginStatus(writer, request, service)
	})).Methods("GET")
}

func oAuthSuccess(auth *strava.AuthorizationResponse, w http.ResponseWriter, r *http.Request, dataService data.IService) {
	log.Debugf("Access Token: %s", auth.AccessToken)

	user, err := dataService.GetStravaUser(auth.Athlete.Id)
	if err != nil {
		log.Error(err)
	}

	if user == nil {
		user = models.NewUser(auth.Athlete.Id)
		user.StravaToken = auth.AccessToken
		user.FirstName = auth.Athlete.FirstName
		user.LastName = auth.Athlete.LastName
		user.ProfileMediumImage = auth.Athlete.ProfileMedium
		user.ProfileImage = auth.Athlete.Profile
		user.Gender = auth.Athlete.Gender
		err = dataService.AddUser(user)
		if err != nil {
			log.Error(err)
		}
	} else {
		user.StravaToken = auth.AccessToken
		user.FirstName = auth.Athlete.FirstName
		user.LastName = auth.Athlete.LastName
		user.ProfileMediumImage = auth.Athlete.ProfileMedium
		user.ProfileImage = auth.Athlete.Profile
		user.Gender = auth.Athlete.Gender
		err = dataService.UpdateUser(*user)
		if err != nil {
			log.Error(err)
		}
	}

	tokenString, err := token.Build(user.Id)
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

func getLoginStatus(w http.ResponseWriter, r *http.Request, dataService data.IService) {
	user, err := token.GetUser(r, dataService)
	if err != nil {
		log.Error(err)
		reply, _ := json.Marshal(false)
		w.WriteHeader(http.StatusNotFound)
		_, err = w.Write(reply)
		if err != nil {
			log.Error(err)
		}
		return
	}

	log.Debugf("name %s %s,  userId: %s", user.FirstName, user.LastName, user.Id)

	reply, _ := json.Marshal(true)
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(reply)
	if err != nil {
		log.Error(err)
	}
}
