package strava

import (
	"time"

	"github.com/cjburchell/go.strava"
)

type IService interface {
	GetStaredSegments() ([]*strava.SegmentSummary, error)
	GetRoutes(athleteId int64) ([]*strava.Route, error)
	GetRoute(routeId int64) (*strava.Route, error)
	GetSegment(segmentId int64) (*strava.SegmentDetailed, error)
}

type service struct {
	client *strava.Client
}

func (s service) GetRoutes(athleteId int64) ([]*strava.Route, error) {
	athletes := strava.NewAthletesService(s.client)
	return athletes.Routes(athleteId).Do()
}

func (s service) GetRoute(routeId int64) (*strava.Route, error) {
	routes := strava.NewRoutesService(s.client)
	return routes.Get(routeId).Do()
}

func (s service) GetRouteStreams(routeId int64, streamTypes []strava.StreamType) (*strava.StreamSet, error) {
	routes := strava.NewRoutesStreamsService(s.client)
	return routes.Get(routeId, streamTypes).Do()
}

func (s service) GetSegmentStream(segmentId int64, streamTypes []strava.StreamType) (*strava.StreamSet, error) {
	routes := strava.NewSegmentStreamsService(s.client)
	return routes.Get(segmentId, streamTypes).Do()
}

func (s service) GetStaredSegments() ([]*strava.SegmentSummary, error) {
	segments := strava.NewSegmentsService(s.client)
	return segments.Starred().PerPage(100).Do()
}

func (s service) SegmentsListEfforts(segmentId int64, athleteId int64, startTime time.Time, endTime time.Time) ([]*strava.SegmentEffortSummary, error) {
	segments := strava.NewSegmentsService(s.client)
	return segments.ListEfforts(segmentId).AthleteId(athleteId).DateRange(startTime, endTime).Do()
}

func (s service) GetSegment(segmentId int64) (*strava.SegmentDetailed, error) {
	segments := strava.NewSegmentsService(s.client)
	return segments.Get(segmentId).Do()
}

func (s service) GetFriends(athleteId int64, page int) ([]*strava.AthleteSummary, error) {
	athlete := strava.NewAthletesService(s.client)
	return athlete.ListFriends(athleteId).Page(page).Do()
}

func NewService(accessToken string) IService {
	client := strava.NewClient(accessToken)
	return &service{client: client}
}
