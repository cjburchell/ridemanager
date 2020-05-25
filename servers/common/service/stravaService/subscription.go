package stravaService

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ISubscription interface {
	Delete() error
}

type subscription struct {
	ID           int    `json:"id"`
	URL          string `json:"url"`
	clientID     int
	clientSecret string
}

func (s subscription) Delete() error {
	client := http.DefaultClient

	req, err := http.NewRequest(
		http.MethodDelete,
		fmt.Sprintf("%s/push_subscriptions/%d?client_id=%d&client_secret=%s",basePath,s.ID, s.clientID, s.clientSecret) ,
		nil)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()


	if resp.StatusCode/100 == 5 {
		return ServerErr
	}

	if resp.StatusCode/100 != 2 {
		var response errorResponse
		contents, _ := ioutil.ReadAll(resp.Body)
		err = json.Unmarshal(contents, &response)
		if err != nil {
			return err
		}

		if len(response.Errors) == 0 {
			return ServerErr
		}

		if response.Errors[0].Resource == "Application" {
			return InvalidCredentialsErr
		}

		return nil
	}

	return nil
}

func getSubscriptions(clientID int, clientSecret string)([]subscription, error)  {
	client := http.DefaultClient

	resp, err := client.Get(fmt.Sprintf("%s/push_subscriptions?client_id=%d&client_secret=%s",basePath, clientID, clientSecret))
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()


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

		return nil, &response
	}

	var response []subscription
	contents, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(contents, &response)
	if err != nil {
		return nil, err
	}

	for i, _ := range response {
		response[i].clientSecret = clientSecret
		response[i].clientID = clientID
	}

	return response, nil
}

func CreateSubscription(clientID int, clientSecret string, callbackURL string)(ISubscription, error) {
	subs, err := getSubscriptions(clientID, clientSecret)
	if err != nil {
		return nil, err
	}

	if len(subs) != 0 {
		for _, sub := range subs {
			if sub.URL == callbackURL {
				// we already have a subscription
				return sub, nil
			}

			err := sub.Delete()
			if err != nil {
				return nil, err
			}
		}
	}

	client := http.DefaultClient

	resp, err := client.PostForm(basePath+"/push_subscriptions",
		url.Values{
			"client_id":     {fmt.Sprintf("%d", clientID)},
			"client_secret": {clientSecret},
			"callback_url":  {callbackURL},
			"verify_token":  {"verify_token"},
		})
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()

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

		return nil, &response
	}

	var response subscription
	contents, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(contents, &response)
	if err != nil {
		return nil, err
	}

	response.clientSecret = clientSecret
	response.clientID = clientID
	return &response, nil
}
