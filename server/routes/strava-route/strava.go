package strava_route

import (
	"encoding/json"
	"github.com/cjburchell/go-uatu"
	"github.com/cjburchell/ridemanager/routes/token"
	"github.com/cjburchell/ridemanager/service/data"
	"github.com/cjburchell/ridemanager/service/data/models"
	"github.com/cjburchell/ridemanager/service/strava"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func Setup(r *mux.Router, service data.IService) {
	dataRoute := r.PathPrefix("/api/v1/strava").Subrouter()

	dataRoute.HandleFunc("/segments/starred", token.ValidateMiddleware(
		func(writer http.ResponseWriter, request *http.Request) {
			handleGetStarredSegments(writer, request, service)
		})).Methods("GET").Queries("page", "{page}", "perPage", "{perPage}")

	dataRoute.HandleFunc("/routes", token.ValidateMiddleware(
		func(writer http.ResponseWriter, request *http.Request) {
			handleGetRoutes(writer, request, service)
		})).Methods("GET").Queries("page", "{page}", "perPage", "{perPage}")

	dataRoute.HandleFunc("/routes/{RouteId}", token.ValidateMiddleware(
		func(writer http.ResponseWriter, request *http.Request) {
			handleGetRoute(writer, request, service)
		})).Methods("GET")

	dataRoute.HandleFunc("/segments/{SegmentId}", token.ValidateMiddleware(
		func(writer http.ResponseWriter, request *http.Request) {
			handleGetSegment(writer, request, service)
		})).Methods("GET")


	dataRoute.HandleFunc("/friends", token.ValidateMiddleware(
		func(writer http.ResponseWriter, request *http.Request) {
			handleGetFriends(writer, request, service)
		})).Methods("GET").Queries("page", "{page}", "perPage", "{perPage}")
}

func handleGetFriends(writer http.ResponseWriter, request *http.Request, service data.IService) {
	user, err := token.GetUser(request, service)
	if err != nil{
		writer.WriteHeader(http.StatusUnauthorized)
		log.Error(err)
		return
	}

	if user == nil{
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	page, err := strconv.Atoi(request.FormValue("page"))
	if err != nil{
		writer.WriteHeader(http.StatusBadRequest)
		log.Error(err)
		return
	}

	perPage, err := strconv.Atoi(request.FormValue("perPage"))
	if err != nil{
		writer.WriteHeader(http.StatusBadRequest)
		log.Error(err)
		return
	}
	stravaService := strava.NewService(user.StravaToken)

	friends, err := stravaService.GetFriends(user.Athlete.StravaAthleteId, page, perPage)
	if err != nil{
		writer.WriteHeader(http.StatusBadRequest)
		log.Error(err)
		return
	}

	friendsList := make([]models.Athlete, len(friends))

	for index, friend := range friends{
		user, err := service.GetStravaUser(friend.Id)
		if err != nil{
			writer.WriteHeader(http.StatusInternalServerError)
			log.Error(err)
			return
		}

		if user == nil{
			user = models.NewUser(friend.Id)
			user.Athlete.Update(*friend)
			err = service.AddUser(user)
			if err != nil {
				writer.WriteHeader(http.StatusInternalServerError)
				log.Error(err)
				return
			}
		} else {
			user.Athlete.Update(*friend)
			err = service.UpdateUser(*user)
			if err != nil {
				writer.WriteHeader(http.StatusInternalServerError)
				log.Error(err)
				return
			}
		}

		friendsList[index] = user.Athlete
	}

	reply, _ := json.Marshal(friendsList)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	_, err = writer.Write(reply)
	if err != nil {
		log.Error(err)
	}
}

func handleGetSegment(writer http.ResponseWriter, request *http.Request, service data.IService) {
	user, err := token.GetUser(request, service)
	if err != nil{
		writer.WriteHeader(http.StatusUnauthorized)
		log.Error(err)
		return
	}

	if user == nil{
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(request)
	segmentId, err := strconv.ParseInt(vars["SegmentId"], 10, 64)
	if err != nil{
		writer.WriteHeader(http.StatusBadRequest)
		log.Error(err)
		return
	}

	stravaService := strava.NewService(user.StravaToken)

	segment, err := stravaService.GetSegment(segmentId)
	if err != nil{
		writer.WriteHeader(http.StatusNotFound)
		log.Error(err)
		return
	}

	if segment == nil{
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	reply, _ := json.Marshal(segment)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	_, err = writer.Write(reply)
	if err != nil {
		log.Error(err)
	}
}

func handleGetRoute(writer http.ResponseWriter, request *http.Request, service data.IService) {
	user, err := token.GetUser(request, service)
	if err != nil{
		writer.WriteHeader(http.StatusUnauthorized)
		log.Error(err)
		return
	}

	if user == nil{
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(request)
	activityIdString := vars["RouteId"]
	activityId, err := strconv.ParseInt(activityIdString, 10, 64)
	if err != nil{
		writer.WriteHeader(http.StatusBadRequest)
		log.Error(err)
		return
	}

	stravaService := strava.NewService(user.StravaToken)

	route, err := stravaService.GetRoute(activityId)
	if err != nil{
		writer.WriteHeader(http.StatusNotFound)
		log.Error(err)
		return
	}

	if route == nil{
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	reply, _ := json.Marshal(route)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	_, err = writer.Write(reply)
	if err != nil {
		log.Error(err)
	}
}

func handleGetRoutes(writer http.ResponseWriter, request *http.Request, service data.IService) {
	user, err := token.GetUser(request, service)
	if err != nil{
		writer.WriteHeader(http.StatusUnauthorized)
		log.Error(err)
		return
	}

	if user == nil{
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}


	page, err := strconv.Atoi(request.FormValue("page"))
	if err != nil{
		writer.WriteHeader(http.StatusBadRequest)
		log.Error(err)
		return
	}

	perPage, err := strconv.Atoi(request.FormValue("perPage"))
	if err != nil{
		writer.WriteHeader(http.StatusBadRequest)
		log.Error(err)
		return
	}

	stravaService := strava.NewService(user.StravaToken)

	routes, err := stravaService.GetRoutes(user.Athlete.StravaAthleteId, page, perPage)
	if err != nil{
		writer.WriteHeader(http.StatusBadRequest)
		log.Error(err)
		return
	}

	reply, _ := json.Marshal(routes)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	_, err = writer.Write(reply)
	if err != nil {
		log.Error(err)
	}
}

func handleGetStarredSegments(writer http.ResponseWriter, request *http.Request, service data.IService) {
	user, err := token.GetUser(request, service)
	if err != nil{
		writer.WriteHeader(http.StatusUnauthorized)
		log.Error(err)
		return
	}

	if user == nil{
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	page, err := strconv.Atoi(request.FormValue("page"))
	if err != nil{
		writer.WriteHeader(http.StatusBadRequest)
		log.Error(err)
		return
	}

	perPage, err := strconv.Atoi(request.FormValue("perPage"))
	if err != nil{
		writer.WriteHeader(http.StatusBadRequest)
		log.Error(err)
		return
	}

	stravaService := strava.NewService(user.StravaToken)

	segments, err := stravaService.GetStaredSegments(page, perPage)
	if err != nil{
		writer.WriteHeader(http.StatusBadRequest)
		log.Error(err)
		return
	}

	reply, _ := json.Marshal(segments)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	_, err = writer.Write(reply)
	if err != nil {
		log.Error(err)
	}
}
