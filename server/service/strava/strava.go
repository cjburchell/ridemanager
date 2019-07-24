package strava

import (
	log "github.com/cjburchell/go-uatu"
	"time"

	"github.com/cjburchell/go.strava"
)

type IService interface {
	GetStaredSegments(page int, perPage int) ([]*strava.SegmentSummary, error)
	GetRoutes(athleteId int64, page int, perPage int) ([]*strava.Route, error)
	GetRoute(routeId int64) (*strava.Route, error)
	GetSegment(segmentId int64) (*strava.SegmentDetailed, error)
}

type service struct {
	client *strava.Client
}

func (s service) GetRoutes(athleteId int64, page int, perPage int) ([]*strava.Route, error) {
	log.Print("Strava GetRoutes")
	athletes := strava.NewAthletesService(s.client)
	return athletes.Routes(athleteId).PerPage(perPage).Page(page+1).Do()
}

func (s service) GetRoute(routeId int64) (*strava.Route, error) {
	log.Print("Strava GetRoute")
	routes := strava.NewRoutesService(s.client)
	return routes.Get(routeId).Do()
}

func (s service) GetRouteStreams(routeId int64, streamTypes []strava.StreamType) (*strava.StreamSet, error) {
	log.Print("Strava GetRouteStreams")
	routes := strava.NewRoutesStreamsService(s.client)
	return routes.Get(routeId, streamTypes).Do()
}

func (s service) GetSegmentStream(segmentId int64, streamTypes []strava.StreamType) (*strava.StreamSet, error) {
	log.Print("Strava GetSegmentStream")
	routes := strava.NewSegmentStreamsService(s.client)
	return routes.Get(segmentId, streamTypes).Do()
}

func (s service) GetStaredSegments(page int, perPage int) ([]*strava.SegmentSummary, error) {
	log.Print("Strava GetStaredSegments")
	segments := strava.NewSegmentsService(s.client)
	return segments.Starred().PerPage(perPage).Page(page + 1).Do()
}

func (s service) SegmentsListEfforts(segmentId int64, athleteId int64, startTime time.Time, endTime time.Time) ([]*strava.SegmentEffortSummary, error) {
	log.Print("Strava SegmentsListEfforts")
	segments := strava.NewSegmentsService(s.client)
	return segments.ListEfforts(segmentId).AthleteId(athleteId).DateRange(startTime, endTime).Do()
}

func (s service) GetSegment(segmentId int64) (*strava.SegmentDetailed, error) {
	log.Print("Strava GetSegment")
	segments := strava.NewSegmentsService(s.client)
	return segments.Get(segmentId).Do()
}

func NewService(accessToken string) IService {
	client := strava.NewClient(accessToken)
	return &service{client: client}
}
