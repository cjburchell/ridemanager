package activityroute

import (
	"encoding/json"
	"net/http"

	"github.com/cjburchell/ridemanager/common/service/stravaService"

	"github.com/cjburchell/ridemanager/common/service/results"

	activityService "github.com/cjburchell/ridemanager/common/service/activity"
	"github.com/cjburchell/ridemanager/common/service/data/models"
	"github.com/cjburchell/tools-go/slice"
	"github.com/cjburchell/uatu-go"

	"github.com/cjburchell/ridemanager/server/routes/token"
	"github.com/cjburchell/ridemanager/common/service/data"
	"github.com/gorilla/mux"
)

type handle struct {
	dataService   data.IService
	authenticator stravaService.Authenticator
	log           log.ILog
}

// Setup the activity route
func Setup(r *mux.Router, service data.IService, tokenValidator token.Validator, authenticator stravaService.Authenticator, logger log.ILog) {

	dataRoute := r.PathPrefix("/api/v1/activity").Subrouter()

	handle := handle{service, authenticator, logger}
	validateWritable := validateWritable{service, logger}

	dataRoute.HandleFunc("/my", tokenValidator.ValidateMiddleware(handle.getMyActivities)).Methods("GET")
	dataRoute.HandleFunc("/public", tokenValidator.ValidateMiddleware(handle.getPublicActivities)).Methods("GET")
	dataRoute.HandleFunc("/joined", tokenValidator.ValidateMiddleware(handle.getJoinedActivities)).Methods("GET")
	dataRoute.HandleFunc("/{ActivityId}", handle.getActivity).Methods("GET")
	dataRoute.HandleFunc("", tokenValidator.ValidateMiddleware(handle.createActivity)).Methods("POST")
	dataRoute.HandleFunc("/{ActivityId}/update", tokenValidator.ValidateMiddleware(validateWritable.Middleware(handle.updateActivityState))).Methods("POST")
	dataRoute.HandleFunc("/{ActivityId}/update/{AthleteId}", tokenValidator.ValidateMiddleware(validateWritable.Middleware(handle.updateActivityParticipantState))).Methods("POST")
	dataRoute.HandleFunc("/{ActivityId}/participant", tokenValidator.ValidateMiddleware(handle.addParticipant)).Methods("POST")
	dataRoute.HandleFunc("/{ActivityId}/participant/{AthleteId}", tokenValidator.ValidateMiddleware(validateWritable.Middleware(handle.removeParticipant))).Methods("DELETE")
	dataRoute.HandleFunc("/{ActivityId}", tokenValidator.ValidateMiddleware(validateWritable.Middleware(handle.updateActivity))).Methods("PATCH")
	dataRoute.HandleFunc("/{ActivityId}", tokenValidator.ValidateMiddleware(validateWritable.Middleware(handle.deleteActivity))).Methods("DELETE")
}

func (h handle) updateResults(activity *models.Activity) error {
	activity.UpdateState()
	if activity.State == models.ActivityStates.Upcoming {
		return nil
	}

	for p := range activity.Participants {

		participant := activity.Participants[p]
		user, err := h.dataService.GetUser(participant.Athlete.Id)
		if err != nil {
			return err
		}

		err = results.UpdateParticipant(participant, activity, stravaService.GetTokenManager(h.authenticator, participant.Athlete.Id, h.dataService, &user.StravaToken))
		if err != nil {
			return err
		}
	}

	results.Update(activity)

	return nil
}

func (h handle) updateActivityState(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	activityID := models.ActivityId(vars["ActivityId"])

	user, err := token.GetUser(request, h.dataService)
	if err != nil || user == nil {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	activity, err := h.dataService.GetActivity(activityID)
	if err != nil {
		h.log.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = h.updateResults(activity)
	if err != nil {
		h.log.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = h.dataService.UpdateActivity(*activity)
	if err != nil {
		h.log.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	reply, _ := json.Marshal(true)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	_, err = writer.Write(reply)
	if err != nil {
		h.log.Error(err)
	}
}

func (h handle) addParticipant(writer http.ResponseWriter, request *http.Request) {
	user, err := token.GetUser(request, h.dataService)
	if err != nil {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	if user == nil {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(request)
	activityID := models.ActivityId(vars["ActivityId"])

	activity, err := h.dataService.GetActivity(activityID)
	if err != nil {
		h.log.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	decoder := json.NewDecoder(request.Body)
	var participant models.Participant
	err = decoder.Decode(&participant)
	if err != nil {
		h.log.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	if user.Athlete.Id != participant.Athlete.Id {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	if activity.Participants == nil {
		activity.Participants = []*models.Participant{&participant}
	} else {
		for _, p := range activity.Participants {
			if p.Athlete.Id == participant.Athlete.Id {
				h.writeSuccess(writer)
				return
			}
		}

		activity.Participants = append(activity.Participants, &participant)
	}

	activity.UpdateState()

	err = results.UpdateParticipant(&participant, activity, stravaService.GetTokenManager(h.authenticator, participant.Athlete.Id, h.dataService, &user.StravaToken))
	if err != nil {
		h.log.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	results.Update(activity)

	err = h.dataService.UpdateActivity(*activity)
	if err != nil {
		h.log.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	h.writeSuccess(writer)
}

func (h handle) writeSuccess(writer http.ResponseWriter) {
	reply, _ := json.Marshal(true)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	_, err := writer.Write(reply)
	if err != nil {
		h.log.Error(err)
	}
}

func (h handle) updateActivityParticipantState(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	activityID := models.ActivityId(vars["ActivityId"])

	user, err := token.GetUser(request, h.dataService)
	if err != nil || user == nil {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	activity, err := h.dataService.GetActivity(activityID)
	if err != nil {
		h.log.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	athleteID := models.AthleteId(vars["AthleteId"])
	participant := activity.FindParticipant(athleteID)
	if participant == nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = results.UpdateParticipant(participant, activity, stravaService.GetTokenManager(h.authenticator, participant.Athlete.Id, h.dataService, &user.StravaToken))
	if err != nil {
		h.log.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	results.Update(activity)

	err = h.dataService.UpdateActivity(*activity)
	if err != nil {
		h.log.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	reply, _ := json.Marshal(true)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	_, err = writer.Write(reply)
	if err != nil {
		h.log.Error(err)
	}
}

func remove(s []*models.Participant, i int) []*models.Participant {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func (h handle) removeParticipant(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	activityID := models.ActivityId(vars["ActivityId"])

	activity, err := h.dataService.GetActivity(activityID)
	if err != nil {
		h.log.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	athleteID := models.AthleteId(vars["AthleteId"])
	index := slice.Index(len(activity.Participants), func(i int) bool {
		return activity.Participants[i].Athlete.Id == athleteID
	})
	if index == -1 {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	activity.Participants = remove(activity.Participants, index)

	err = h.dataService.UpdateActivity(*activity)
	if err != nil {
		h.log.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	reply, _ := json.Marshal(true)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	_, err = writer.Write(reply)
	if err != nil {
		h.log.Error(err)
	}
}

func (h handle) deleteActivity(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	activityID := models.ActivityId(vars["ActivityId"])

	err := h.dataService.DeleteActivity(activityID)
	if err != nil {
		h.log.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
}

func (h handle) updateActivity(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var changedActivity models.Activity
	err := decoder.Decode(&changedActivity)
	if err != nil {
		h.log.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	changedActivity.UpdateState()

	err = h.dataService.UpdateActivity(changedActivity)
	if err != nil {
		h.log.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
}

func (h handle) createActivity(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var newActivity models.Activity
	err := decoder.Decode(&newActivity)
	if err != nil {
		h.log.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	newActivity.UpdateState()

	id, err := h.dataService.AddActivity(&newActivity)
	if err != nil {
		h.log.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	reply, _ := json.Marshal(id)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	_, err = writer.Write(reply)
	if err != nil {
		h.log.Error(err)
	}
}

func (h handle) getMyActivities(writer http.ResponseWriter, req *http.Request) {

	user, err := token.GetUser(req, h.dataService)
	if err != nil || user == nil {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	activities, err := h.dataService.GetOwnedActivities(user.Athlete.Id)
	if err != nil {
		h.log.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = activityService.UpdateAll(activities, h.dataService, false, h.authenticator)
	if err != nil {
		h.log.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	reply, _ := json.Marshal(activities)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	_, err = writer.Write(reply)
	if err != nil {
		h.log.Error(err)
	}
}

func (h handle) getJoinedActivities(writer http.ResponseWriter, req *http.Request) {

	user, err := token.GetUser(req, h.dataService)
	if err != nil || user == nil {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	activities, err := h.dataService.GetAthleteActivities(user.Athlete.Id)
	if err != nil {
		h.log.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = activityService.UpdateAll(activities, h.dataService, false, h.authenticator)
	if err != nil {
		h.log.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	reply, _ := json.Marshal(activities)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	_, err = writer.Write(reply)
	if err != nil {
		h.log.Error(err)
	}
}

func (h handle) getPublicActivities(writer http.ResponseWriter, _ *http.Request) {
	activities, err := h.dataService.GetActivitiesByPrivacy(models.Privacy.Public)
	if err != nil {
		h.log.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = activityService.UpdateAll(activities, h.dataService, false, h.authenticator)
	if err != nil {
		h.log.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	reply, _ := json.Marshal(activities)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	_, err = writer.Write(reply)
	if err != nil {
		h.log.Error(err)
	}
}

func (h handle) getActivity(writer http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)
	activityID := models.ActivityId(vars["ActivityId"])

	foundActivity, err := h.dataService.GetActivity(activityID)
	if err != nil {
		h.log.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	if foundActivity.UpdateState() {
		err = h.dataService.UpdateActivity(*foundActivity)
		if err != nil {
			h.log.Error(err)
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	reply, _ := json.Marshal(foundActivity)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	_, err = writer.Write(reply)
	if err != nil {
		h.log.Error(err)
	}

}
