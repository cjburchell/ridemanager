package activity_route

import (
	"encoding/json"
	"github.com/cjburchell/go-uatu"
	activityservice "github.com/cjburchell/ridemanager/service/activity"
	"github.com/cjburchell/ridemanager/service/data/models"
	"github.com/cjburchell/tools-go/slice"
	"net/http"

	"github.com/cjburchell/ridemanager/routes/token"
	"github.com/cjburchell/ridemanager/service/data"
	"github.com/gorilla/mux"
)

func Setup(r *mux.Router, service data.IService) {

	dataRoute := r.PathPrefix("/api/v1/activity").Subrouter()

	dataRoute.HandleFunc("/my", token.ValidateMiddleware(
		func(writer http.ResponseWriter, request *http.Request) {
			handleGetMyActivities(writer, request, service)
		})).Methods("GET")

	dataRoute.HandleFunc("/public", token.ValidateMiddleware(
		func(writer http.ResponseWriter, request *http.Request) {
			handleGetPublicActivities(writer, request, service)
		})).Methods("GET")

	dataRoute.HandleFunc("/joined", token.ValidateMiddleware(
		func(writer http.ResponseWriter, request *http.Request) {
			handleGetJoinedActivities(writer, request, service)
		})).Methods("GET")

	dataRoute.HandleFunc("/{ActivityId}",
		func(writer http.ResponseWriter, request *http.Request) {
			handleGetActivity(writer, request, service)
		}).Methods("GET")

	dataRoute.HandleFunc("", token.ValidateMiddleware(
		func(writer http.ResponseWriter, request *http.Request) {
			handleCreateActivity(writer, request, service)
		})).Methods("POST")

	dataRoute.HandleFunc("{ActivityId}/update", token.ValidateMiddleware(validateWritableAccessMiddleware(
		func(writer http.ResponseWriter, request *http.Request) {
			handleUpdateActivityState(writer, request, service)
		}, service))).Methods("POST")

	dataRoute.HandleFunc("{ActivityId}/update/{AthleteId}", token.ValidateMiddleware(validateWritableAccessMiddleware(
		func(writer http.ResponseWriter, request *http.Request) {
			handleUpdateActivityParticipantState(writer, request, service)
		}, service))).Methods("POST")

	dataRoute.HandleFunc("/{ActivityId}/participant", token.ValidateMiddleware(
		func(writer http.ResponseWriter, request *http.Request) {
			handleAddParticipant(writer, request, service)
		})).Methods("POST")

	dataRoute.HandleFunc("/{ActivityId}/participant/{AthleteId}", token.ValidateMiddleware(validateWritableAccessMiddleware(
		func(writer http.ResponseWriter, request *http.Request) {
			handleRemoveParticipant(writer, request, service)
		}, service))).Methods("DELETE")

	dataRoute.HandleFunc("/{ActivityId}", token.ValidateMiddleware(validateWritableAccessMiddleware(
		func(writer http.ResponseWriter, request *http.Request) {
			handleUpdateActivity(writer, request, service)
		}, service))).Methods("PATCH")

	dataRoute.HandleFunc("/{ActivityId}", token.ValidateMiddleware(validateWritableAccessMiddleware(
		func(writer http.ResponseWriter, request *http.Request) {
			handleDeleteActivity(writer, request, service)
		}, service))).Methods("DELETE")
}

func handleUpdateActivityState(writer http.ResponseWriter, request *http.Request, service data.IService) {
	writer.WriteHeader(http.StatusNotImplemented)
}

func handleAddParticipant(writer http.ResponseWriter, request *http.Request, service data.IService) {
	user, err := token.GetUser(request, service)
	if err != nil{
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	if user == nil{
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(request)
	activityId := models.ActivityId(vars["ActivityId"])

	activity, err := service.GetActivity(activityId)
	if err != nil {
		log.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	decoder := json.NewDecoder(request.Body)
	var participant models.Participant
	err = decoder.Decode(&participant)
	if err != nil {
		log.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	if user.Role != models.AdminRole && user.Athlete.Id != activity.Owner.Id && user.Athlete.Id != participant.Athlete.Id {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	if activity.Participants == nil{
		activity.Participants = []models.Participant{participant}
	} else {
		activity.Participants = append(activity.Participants, participant)
	}

	_ = activityservice.UpdateState(activity, nil)

	err = service.UpdateActivity(*activity)
	if err != nil {
		log.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	reply, _ := json.Marshal(true)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	_, err = writer.Write(reply)
	if err != nil {
		log.Error(err)
	}
}

func handleUpdateActivityParticipantState(writer http.ResponseWriter, request *http.Request, service data.IService) {
	writer.WriteHeader(http.StatusNotImplemented)
}

func remove(s []models.Participant, i int) []models.Participant {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func handleRemoveParticipant(writer http.ResponseWriter, request *http.Request, service data.IService) {
	vars := mux.Vars(request)
	activityId := models.ActivityId(vars["ActivityId"])

	activity, err := service.GetActivity(activityId)
	if err != nil {
		log.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	athleteId := models.AthleteId(vars["AthleteId"])
	index := slice.Index(len(activity.Participants), func(i int) bool {
		return activity.Participants[i].Athlete.Id == athleteId
	})
	if index == -1{
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	activity.Participants = remove(activity.Participants, index)

	err = service.UpdateActivity(*activity)
	if err != nil {
		log.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	reply, _ := json.Marshal(true)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	_, err = writer.Write(reply)
	if err != nil {
		log.Error(err)
	}
}

func validateWritableAccessMiddleware(next http.HandlerFunc, service data.IService) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		 user, err := token.GetUser(req, service)
		 if err != nil{
			 w.WriteHeader(http.StatusUnauthorized)
			 return
		 }

		if user == nil{
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// check user role
		if user.Role == models.AdminRole{
			next(w, req)
			return
		}

		 // check activity owner
		vars := mux.Vars(req)
		activityId := models.ActivityId(vars["ActivityId"])

		activity, err := service.GetActivity(activityId)
		if err != nil {
			log.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if activity.Owner.Id != user.Athlete.Id{
			activityId := models.AthleteId(vars["AthleteId"])
			if activityId != user.Athlete.Id{
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
		}

		next(w, req)
	})
}

func handleDeleteActivity(writer http.ResponseWriter, request *http.Request, service data.IService) {
	vars := mux.Vars(request)
	activityId := models.ActivityId(vars["ActivityId"])

	err := service.DeleteActivity(activityId)
	if err != nil {
		log.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
}

func handleUpdateActivity(writer http.ResponseWriter, request *http.Request, service data.IService) {
	decoder := json.NewDecoder(request.Body)
	var changedActivity models.Activity
	err := decoder.Decode(&changedActivity)
	if err != nil {
		log.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	_ = activityservice.UpdateState(&changedActivity, nil)

	err = service.UpdateActivity(changedActivity)
	if err != nil {
		log.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
}

func handleCreateActivity(writer http.ResponseWriter, request *http.Request, service data.IService) {
	decoder := json.NewDecoder(request.Body)
	var newActivity models.Activity
	err := decoder.Decode(&newActivity)
	if err != nil {
		log.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	_ = activityservice.UpdateState(&newActivity, nil)

	id, err := service.AddActivity(&newActivity)
	if err != nil {
		log.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	reply, _ := json.Marshal(id)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	_, err = writer.Write(reply)
	if err != nil {
		log.Error(err)
	}
}

func handleGetMyActivities(writer http.ResponseWriter, req *http.Request, service data.IService) {

	user, err := token.GetUser(req, service)
	if err != nil{
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	if user == nil{
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	activities, err := service.GetOwnedActivities(user.Athlete.Id)
	if err != nil {
		log.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = activityservice.UpdateAll(activities, service)
	if err != nil {
		log.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	reply, _ := json.Marshal(activities)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	_, err = writer.Write(reply)
	if err != nil {
		log.Error(err)
	}
}

func handleGetJoinedActivities(writer http.ResponseWriter, req *http.Request, service data.IService) {

	user, err := token.GetUser(req, service)
	if err != nil{
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	if user == nil{
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	activities, err := service.GetAthleteActivities(user.Athlete.Id)
	if err != nil {
		log.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	err := activityservice.UpdateAll(activities, service)
	if err != nil {
		log.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	reply, _ := json.Marshal(activities)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	_, err = writer.Write(reply)
	if err != nil {
		log.Error(err)
	}
}

func handleGetPublicActivities(writer http.ResponseWriter, _ *http.Request, service data.IService) {
	activities, err := service.GetActivitiesByPrivacy(models.Privacy.Public)
	if err != nil {
		log.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = activityservice.UpdateAll(activities, service)
	if err != nil {
		log.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	reply, _ := json.Marshal(activities)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	_, err = writer.Write(reply)
	if err != nil {
		log.Error(err)
	}
}

func handleGetActivity(writer http.ResponseWriter, request *http.Request, service data.IService) {

	vars := mux.Vars(request)
	activityId := models.ActivityId(vars["ActivityId"])

	foundActivity, err := service.GetActivity(activityId)
	if err != nil {
		log.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = activityservice.UpdateState(foundActivity, service)
	if err != nil {
		log.Error(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	reply, _ := json.Marshal(foundActivity)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	_, err = writer.Write(reply)
	if err != nil {
		log.Error(err)
	}

}
