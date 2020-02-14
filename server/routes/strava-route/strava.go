package strava_route

import (
	"encoding/json"
	"github.com/cjburchell/ridemanager/service/data/models"
	"net/http"
	"strconv"

	"github.com/cjburchell/go-uatu"
	"github.com/cjburchell/ridemanager/routes/token"
	"github.com/cjburchell/ridemanager/service/data"
	"github.com/cjburchell/ridemanager/service/stravaService"
	"github.com/gorilla/mux"
)

type handler struct {
	log           log.ILog
	service       data.IService
	authenticator stravaService.Authenticator
}

func Setup(r *mux.Router, service data.IService, validator token.Validator, authenticator stravaService.Authenticator, logger log.ILog) {
	dataRoute := r.PathPrefix("/api/v1/strava").Subrouter()

	handle := handler{logger, service, authenticator}
	dataRoute.HandleFunc("/segments/starred", validator.ValidateMiddleware(handle.getStarredSegments)).Methods("GET").Queries("page", "{page}", "perPage", "{perPage}")
	dataRoute.HandleFunc("/routes", validator.ValidateMiddleware(handle.getRoutes)).Methods("GET").Queries("page", "{page}", "perPage", "{perPage}")
	dataRoute.HandleFunc("/routes/{RouteId}", validator.ValidateMiddleware(handle.getRoute)).Methods("GET")
	dataRoute.HandleFunc("/segments/{SegmentId}", validator.ValidateMiddleware(handle.getSegment)).Methods("GET")
	dataRoute.HandleFunc("/routes/{RouteId}/elevation", validator.ValidateMiddleware(handle.getRouteElevation)).Methods("GET")
	dataRoute.HandleFunc("/segments/{SegmentId}/elevation", validator.ValidateMiddleware(handle.getSegmentElevation)).Methods("GET")
}

func (h handler) getSegment(writer http.ResponseWriter, request *http.Request) {
	user, err := token.GetUser(request, h.service)
	if err != nil {
		writer.WriteHeader(http.StatusUnauthorized)
		h.log.Error(err)
		return
	}

	if user == nil {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(request)
	segmentId, err := strconv.ParseInt(vars["SegmentId"], 10, 64)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		h.log.Error(err)
		return
	}

	strava := stravaService.NewService(stravaService.GetTokenManager(h.authenticator, user.Athlete.Id, h.service, &user.StravaToken))

	segment, err := strava.GetSegment(segmentId)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		h.log.Error(err)
		return
	}

	if segment == nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	reply, _ := json.Marshal(segment)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	_, err = writer.Write(reply)
	if err != nil {
		h.log.Error(err)
	}
}

func (h handler) getRoute(writer http.ResponseWriter, request *http.Request) {
	user, err := token.GetUser(request, h.service)
	if err != nil {
		writer.WriteHeader(http.StatusUnauthorized)
		h.log.Error(err)
		return
	}

	if user == nil {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(request)
	activityIdString := vars["RouteId"]
	activityId, err := strconv.ParseInt(activityIdString, 10, 32)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		h.log.Error(err)
		return
	}

	strava := stravaService.NewService(stravaService.GetTokenManager(h.authenticator, user.Athlete.Id, h.service, &user.StravaToken))

	route, err := strava.GetRoute(int32(activityId))
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		h.log.Error(err)
		return
	}

	if route == nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	reply, _ := json.Marshal(route)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	_, err = writer.Write(reply)
	if err != nil {
		h.log.Error(err)
	}
}

func (h handler) getRoutes(writer http.ResponseWriter, request *http.Request) {
	user, err := token.GetUser(request, h.service)
	if err != nil {
		writer.WriteHeader(http.StatusUnauthorized)
		h.log.Error(err)
		return
	}

	if user == nil {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	page, err := strconv.Atoi(request.FormValue("page"))
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		h.log.Error(err)
		return
	}

	perPage, err := strconv.Atoi(request.FormValue("perPage"))
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		h.log.Error(err)
		return
	}

	strava := stravaService.NewService(stravaService.GetTokenManager(h.authenticator, user.Athlete.Id, h.service, &user.StravaToken))

	routes, err := strava.GetRoutes(user.Athlete.StravaAthleteId, int32(page), int32(perPage))
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		h.log.Error(err)
		return
	}

	reply, _ := json.Marshal(routes)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	_, err = writer.Write(reply)
	if err != nil {
		h.log.Error(err)
	}
}

func (h handler) getStarredSegments(writer http.ResponseWriter, request *http.Request) {
	user, err := token.GetUser(request, h.service)
	if err != nil {
		writer.WriteHeader(http.StatusUnauthorized)
		h.log.Error(err)
		return
	}

	if user == nil {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	page, err := strconv.Atoi(request.FormValue("page"))
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		h.log.Error(err)
		return
	}

	perPage, err := strconv.Atoi(request.FormValue("perPage"))
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		h.log.Error(err)
		return
	}

	strava := stravaService.NewService(stravaService.GetTokenManager(h.authenticator, user.Athlete.Id, h.service, &user.StravaToken))

	segments, err := strava.GetStaredSegments(int32(page), int32(perPage))
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		h.log.Error(err)
		return
	}

	reply, _ := json.Marshal(segments)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	_, err = writer.Write(reply)
	if err != nil {
		h.log.Error(err)
	}
}

func (h handler) getSegmentElevation(writer http.ResponseWriter, request *http.Request) {
	user, err := token.GetUser(request, h.service)
	if err != nil {
		writer.WriteHeader(http.StatusUnauthorized)
		h.log.Error(err)
		return
	}

	if user == nil {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(request)
	segmentId, err := strconv.ParseInt(vars["SegmentId"], 10, 64)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		h.log.Error(err)
		return
	}

	strava := stravaService.NewService(stravaService.GetTokenManager(h.authenticator, user.Athlete.Id, h.service, &user.StravaToken))

	streamSet, err := strava.GetSegmentStream(segmentId, []string{"distance", "altitude"})
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		h.log.Error(err)
		return
	}

	if streamSet == nil || streamSet.Altitude == nil || streamSet.Distance == nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	elevation := make( []models.Elevation, len(streamSet.Distance.Data))
	for i:=0; i<len(streamSet.Distance.Data); i++{
		elevation[i].X = streamSet.Distance.Data[i]
		elevation[i].Y = streamSet.Altitude.Data[i]
	}

	reply, _ := json.Marshal(elevation)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	_, err = writer.Write(reply)
	if err != nil {
		h.log.Error(err)
	}
}

func (h handler) getRouteElevation(writer http.ResponseWriter, request *http.Request) {
	h.log.Debug("getRouteElevation")
	user, err := token.GetUser(request, h.service)
	if err != nil {
		writer.WriteHeader(http.StatusUnauthorized)
		h.log.Error(err)
		return
	}

	if user == nil {
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(request)
	activityIdString := vars["RouteId"]
	activityId, err := strconv.ParseInt(activityIdString, 10, 32)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		h.log.Error(err)
		return
	}

	strava := stravaService.NewService(stravaService.GetTokenManager(h.authenticator, user.Athlete.Id, h.service, &user.StravaToken))

	streamSet, err := strava.GetRouteStreams(activityId, []string{"distance", "altitude"})
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		h.log.Error(err)
		return
	}

	if streamSet == nil || streamSet.Altitude == nil || streamSet.Distance == nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	elevation := make( []models.Elevation, len(streamSet.Distance.Data))
	for i:=0; i<len(streamSet.Distance.Data); i++{
		elevation[i].X = streamSet.Distance.Data[i]
		elevation[i].Y = streamSet.Altitude.Data[i]
	}

	reply, _ := json.Marshal(elevation)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	_, err = writer.Write(reply)
	if err != nil {
		h.log.Error(err)
	}
}
