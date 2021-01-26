package login_route

import (
	"encoding/json"
	"net/http"

	"github.com/cjburchell/ridemanager/server/routes/token"
	"github.com/cjburchell/ridemanager/common/service/data"
	"github.com/cjburchell/ridemanager/common/service/data/models"
	"github.com/cjburchell/ridemanager/common/service/stravaService"
	"github.com/cjburchell/uatu-go"
	"github.com/gorilla/mux"
)

type handler struct {
	log           log.ILog
	authenticator stravaService.Authenticator
	dataService   data.IService
	tokenBuilder  token.Builder
}

// SetupDataRoute setup the route
func Setup(r *mux.Router, service data.IService, tokenValidator token.Validator, tokenBuilder token.Builder, authenticator stravaService.Authenticator, logger log.ILog) {
	dataRoute := r.PathPrefix("/api/v1/login").Subrouter()
	handle := handler{logger, authenticator, service, tokenBuilder}
	dataRoute.HandleFunc("/validate", handle.validate).Methods("GET").Queries("code", "{code}", "scope", "{scope}", "state", "{state}")
	dataRoute.HandleFunc("/status", tokenValidator.ValidateMiddleware(handle.loginStatus)).Methods("GET")
}

func (h handler) validate(w http.ResponseWriter, r *http.Request) {
	// user denied authorization
	if r.FormValue("error") == "access_denied" {
		h.failure(stravaService.AuthorizationDeniedErr, w, r)
		return
	}

	resp, err := h.authenticator.Authorize(r.FormValue("code"))
	if err != nil {
		h.failure(err, w, r)
		return
	}

	resp.State = r.FormValue("state")

	h.success(resp, w, r)
}

func (h handler) success(auth *stravaService.AuthorizationResponse, w http.ResponseWriter, r *http.Request) {
	h.log.Debugf("Access Token: %s", auth.AccessToken)

	user, err := h.dataService.GetStravaUser(auth.Athlete.Id)
	if err != nil {
		h.log.Error(err)
	}

	if user == nil {
		user = models.NewUser(auth.Athlete.Id)
		user.StravaToken = auth.Token
		user.Athlete.Update(auth.Athlete)
		err = h.dataService.AddUser(user)
		if err != nil {
			h.log.Error(err)
		}
	} else {
		if user.StravaToken.AccessToken != auth.AccessToken {
			h.log.Printf("Token Changed for %s", user.Athlete.Name)
		}

		user.StravaToken = auth.Token
		user.Athlete.Update(auth.Athlete)
		err = h.dataService.UpdateUser(*user)
		if err != nil {
			h.log.Error(err)
		}
	}

	tokenString, err := h.tokenBuilder.BuildToken(user.Athlete.Id)
	if err != nil {
		h.log.Error(err)
	}

	http.Redirect(w, r, "/token?token="+tokenString, http.StatusFound)
}

func (h handler) failure(err error, w http.ResponseWriter, r *http.Request) {

	message := "Authorization Failure:\n"

	// some standard error checking
	if err == stravaService.AuthorizationDeniedErr {
		message += "The user clicked the 'Do not Authorize' button on the previous page."
	} else if err == stravaService.InvalidCredentialsErr {
		message += "You provided an incorrect client_id or client_secret.\nDid you remember to set them at the beginning of this file?"
	} else if err == stravaService.InvalidCodeErr {
		message += "The temporary token was not recognized, this shouldn't happen normally"
	} else if err == stravaService.ServerErr {
		message += "There was some sort of server error, try again to see if the problem continues"
	}

	h.log.Debug(message)
	h.log.Error(err)

	http.Redirect(w, r, "/login", http.StatusFound)
}

func (h handler) loginStatus(w http.ResponseWriter, r *http.Request) {
	user, err := token.GetUser(r, h.dataService)
	if err != nil || user == nil {
		if err != nil {
			h.log.Error(err)
		}

		reply, _ := json.Marshal(false)
		w.WriteHeader(http.StatusNotFound)
		_, err = w.Write(reply)
		if err != nil {
			h.log.Error(err)
		}
		return
	}

	h.log.Debugf("name %s,  userId: %s", user.Athlete.Name, user.Athlete.Id)

	reply, _ := json.Marshal(true)
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(reply)
	if err != nil {
		h.log.Error(err)
	}
}
