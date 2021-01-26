package webhook

import (
	"encoding/json"
	"fmt"
	"github.com/cjburchell/ridemanager/common/service/data"
	"github.com/cjburchell/ridemanager/common/service/stravaService"
	log "github.com/cjburchell/uatu-go"
	"github.com/gorilla/mux"
	"net/http"
)

type handler struct {
	logger log.ILog
	authenticator stravaService.Authenticator
	service data.IService	
	webHookToken string
}

type Request struct {
	ObjectType string `json:"object_type"`
	ObjectID int64 `json:"object_id"`
	AspectType string `json:"aspect_type"`
	OwnerID	int64 `json:"owner_id"`
	SubscriptionID int `json:"subscription_id"`
	EventTime int64 `json:"event_time"`
	Updates json.RawMessage `json:"updates"`
}

func (h handler) webHook(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var requestData Request
	err := decoder.Decode(&requestData)
	if err != nil {
		h.logger.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

 	writer.WriteHeader(http.StatusOK)
}

func (h handler) challenge(writer http.ResponseWriter, request *http.Request) {
	mode := request.FormValue("mode")
	token := request.FormValue("verify_token")
	challenge := request.FormValue("challenge")
	if mode != "subscribe" || token != h.webHookToken {
		writer.WriteHeader(http.StatusForbidden)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	_, err := writer.Write([]byte(fmt.Sprintf(`{"hub.challenge":%s}`, challenge)))
	if err != nil {
		h.logger.Error(err)
	}

	h.logger.Print("Web Hook Verified")
}

// SetupDataRoute setup the route
func Setup(r *mux.Router, service data.IService, authenticator stravaService.Authenticator, logger log.ILog, webHookToken string) {
	dataRoute := r.PathPrefix("/api/v1/webhook").Subrouter()
	handle := handler{logger, authenticator, service, webHookToken}
	dataRoute.HandleFunc("/push", handle.webHook).Methods("POST")
	dataRoute.HandleFunc("/push", handle.challenge).Methods("GET").Queries("hub.mode", "{mode}", "hub.verify_token", "{verify_token}", "hub.challenge", "{challenge}")
}