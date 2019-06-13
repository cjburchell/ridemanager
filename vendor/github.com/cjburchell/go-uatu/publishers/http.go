package publishers

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/cjburchell/go-uatu"
)

type HttpSettings struct {
	Address string
	Token   string
}

type httpPublisher struct {
	restClient *http.Client
	settings   HttpSettings
}

func (publisher httpPublisher) Publish(messageBites []byte) error {
	req, err := http.NewRequest("POST", publisher.settings.Address, bytes.NewBuffer(messageBites))
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", fmt.Sprintf("APIKEY %s", publisher.settings.Token))
	req.Header.Add("Content-Type", "application/json")

	resp, err := publisher.restClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("Unable to send log to %s(%d)", publisher.settings.Address, resp.StatusCode)
	}

	return nil
}

func SetupHttp(newSettings HttpSettings) log.Publisher {
	restClient := &http.Client{}
	return httpPublisher{restClient: restClient, settings: newSettings}
}
