package stravaService

import (
	"context"

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

func (s service) GetRouteStreams(routeId int64, streamTypes []string) (*strava.StreamSet, error) {
	ctx, err := s.getContext()
	if err != nil {
		return nil, err
	}

	result, _, err := s.client.StreamsApi.GetRouteStreams(ctx, routeId, streamTypes, true)
	return &result, err
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
