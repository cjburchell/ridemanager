package stravaService

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/antihax/optional"
	"github.com/cjburchell/strava-go"
)

type IService interface {
	GetStaredSegments(page int32, perPage int32) ([]strava.SummarySegment, error)
	GetRoutes(athleteId int32, page int32, perPage int32) ([]strava.Route, error)
	GetRoute(routeId int32) (*strava.Route, error)
	GetSegment(segmentId int64) (*strava.DetailedSegment, error)
	SegmentsListEfforts(segmentId int32, page int32, perPage int32) ([]strava.DetailedSegmentEffort, error)
	GetRouteStreams(routeId int64, streamTypes []string) (*strava.StreamSet, error)
	GetSegmentStream(segmentId int64, streamTypes []string) (*strava.StreamSet, error)
}

type service struct {
	client *strava.APIClient
	token  TokenManager
}

func (s service) GetRoutes(athleteId int32, page int32, perPage int32) ([]strava.Route, error) {
	ctx, err := s.getContext()
	if err != nil {
		return nil, err
	}

	routes, _, err := s.client.RoutesApi.GetRoutesByAthleteId(ctx, athleteId, &strava.GetRoutesByAthleteIdOpts{Page: optional.NewInt32(page + 1), PerPage: optional.NewInt32(perPage)})
	return routes, err
}

func (s service) getContext() (context.Context, error) {
	ctx := context.Background()
	token, err := s.token.Get()
	if err != nil {
		return nil, err
	}

	ctx = context.WithValue(ctx, strava.ContextAccessToken, token.AccessToken)
	return ctx, nil
}

func (s service) GetRoute(routeId int32) (*strava.Route, error) {
	ctx, err := s.getContext()
	if err != nil {
		return nil, err
	}

	result, _, err := s.client.RoutesApi.GetRouteById(ctx, routeId)
	return &result, err
}

func (s service) GetRouteStreams(routeId int64) (*strava.StreamSet, error) {
	ctx, err := s.getContext()
	if err != nil {
		return nil, err
	}

	result, _, err := s.client.StreamsApi.GetRouteStreams(ctx, routeId)
	return &result, err
}

func GetRouteStreams(a *strava.StreamsApiService, ctx context.Context, id int64) (strava.StreamSet, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue strava.StreamSet
	)

	// create path and map variables
	cfg := strava.NewConfiguration();
	localVarPath := cfg.BasePath + "/routes/{id}/streams"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	http.

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := strava.selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := strava.selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := strava.GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		if localVarHttpResponse.StatusCode == 200 {
			var v strava.StreamSet
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		if localVarHttpResponse.StatusCode == 0 {
			var v strava.Fault
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

func (s service) GetSegmentStream(segmentId int64, streamTypes []string) (*strava.StreamSet, error) {
	ctx, err := s.getContext()
	if err != nil {
		return nil, err
	}

	result, _, err := s.client.StreamsApi.GetSegmentStreams(ctx, segmentId, streamTypes, true)
	return &result, err
}

func (s service) GetStaredSegments(page int32, perPage int32) ([]strava.SummarySegment, error) {
	ctx, err := s.getContext()
	if err != nil {
		return nil, err
	}

	result, _, err := s.client.SegmentsApi.GetLoggedInAthleteStarredSegments(ctx, &strava.GetLoggedInAthleteStarredSegmentsOpts{
		Page:    optional.NewInt32(page + 1),
		PerPage: optional.NewInt32(perPage),
	})
	return result, err
}

func (s service) SegmentsListEfforts(segmentId int32, page int32, perPage int32) ([]strava.DetailedSegmentEffort, error) {
	ctx, err := s.getContext()
	if err != nil {
		return nil, err
	}

	result, _, err := s.client.SegmentEffortsApi.GetEffortsBySegmentId(ctx, segmentId, &strava.GetEffortsBySegmentIdOpts{
		Page:    optional.NewInt32(page + 1),
		PerPage: optional.NewInt32(perPage),
	})
	return result, err
}

func (s service) GetSegment(segmentId int64) (*strava.DetailedSegment, error) {
	ctx, err := s.getContext()
	if err != nil {
		return nil, err
	}

	result, _, err := s.client.SegmentsApi.GetSegmentById(ctx, segmentId)
	return &result, err
}

func NewService(token TokenManager) IService {
	client := strava.NewAPIClient(strava.NewConfiguration())
	return &service{client: client, token: token}
}
