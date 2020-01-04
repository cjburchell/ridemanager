package stravaService

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/cjburchell/ridemanager/service/data/models"
	"github.com/cjburchell/strava-go"
)

const basePath = "https://www.strava.com/api/v3"

// AuthorizationResponse is returned as a result of the token exchange
type AuthorizationResponse struct {
	models.Token
	Athlete strava.SummaryAthlete `json:"athlete"`
	State   string
}

type Authenticator interface {
	Renew(refreshToken string) (*models.Token, error)
	Authorize(code string) (*AuthorizationResponse, error)
}

func GetAuthenticator(clientId int, clientSecret string) Authenticator {
	return authenticator{clientId, clientSecret}
}

type authenticator struct {
	ClientId     int
	ClientSecret string
}

// returned during oauth if there was a user caused problem
// such as user did not grant access or the id/secret was invalid
type AuthError struct {
	message string
}

func (e *AuthError) Error() string {
	return e.message
}

var (
	AuthorizationDeniedErr = &AuthError{"authorization denied by user"}
	InvalidCredentialsErr  = &AuthError{"invalid client_id or client_secret"}
	InvalidCodeErr         = &AuthError{"unrecognized code"}
	ServerErr              = &AuthError{"server error"}
)

type errorDetailed struct {
	Resource string `json:"resource"`
	Field    string `json:"field"`
	Code     string `json:"code"`
}

type errorResponse struct {
	Message string           `json:"message"`
	Errors  []*errorDetailed `json:"errors"`
}

func (e errorResponse) Error() string {
	b, _ := json.Marshal(e)
	return string(b)
}

func (auth authenticator) Authorize(code string) (*AuthorizationResponse, error) {
	// make sure a code was passed
	if code == "" {
		return nil, InvalidCodeErr
	}

	client := http.DefaultClient

	resp, err := client.PostForm(basePath+"/oauth/token",
		url.Values{
			"client_id":     {fmt.Sprintf("%d", auth.ClientId)},
			"client_secret": {auth.ClientSecret},
			"code":          {code},
			"grant_type":    {"authorization_code"},
		})

	// this was a poor request, maybe strava servers down?
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	// check status code, could be 500, or most likely the client_secret is incorrect
	if resp.StatusCode/100 == 5 {
		return nil, ServerErr
	}

	if resp.StatusCode/100 != 2 {
		var response errorResponse
		contents, _ := ioutil.ReadAll(resp.Body)
		err = json.Unmarshal(contents, &response)
		if err != nil {
			return nil, err
		}

		if len(response.Errors) == 0 {
			return nil, ServerErr
		}

		if response.Errors[0].Resource == "Application" {
			return nil, InvalidCredentialsErr
		}

		if response.Errors[0].Resource == "RequestToken" {
			return nil, InvalidCodeErr
		}

		return nil, &response
	}

	var response AuthorizationResponse
	contents, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(contents, &response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (auth authenticator) Renew(refreshToken string) (*models.Token, error) {
	// make sure a code was passed
	if refreshToken == "" {
		return nil, InvalidCodeErr
	}

	client := http.DefaultClient

	resp, err := client.PostForm(basePath+"/oauth/token",
		url.Values{
			"client_id":     {fmt.Sprintf("%d", auth.ClientId)},
			"client_secret": {auth.ClientSecret},
			"refresh_token": {refreshToken},
			"grant_type":    {"refresh_token"},
		})

	// this was a poor request, maybe strava servers down?
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

	// check status code, could be 500, or most likely the client_secret is incorrect
	if resp.StatusCode/100 == 5 {
		return nil, ServerErr
	}

	if resp.StatusCode/100 != 2 {
		var response errorResponse
		contents, _ := ioutil.ReadAll(resp.Body)
		err = json.Unmarshal(contents, &response)
		if err != nil {
			return nil, err
		}

		if len(response.Errors) == 0 {
			return nil, ServerErr
		}

		if response.Errors[0].Resource == "Application" {
			return nil, InvalidCredentialsErr
		}

		if response.Errors[0].Resource == "RequestToken" {
			return nil, InvalidCodeErr
		}

		return nil, &response
	}

	var response models.Token
	contents, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(contents, &response)

	if err != nil {
		return nil, err
	}

	return &response, nil
}
