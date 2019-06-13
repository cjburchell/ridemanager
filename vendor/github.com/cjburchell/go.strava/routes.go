package strava

import (
	"encoding/json"
	"fmt"
)

type Map struct {
	Id              int64  `json:"id"`
	Polyline        string `json:"polyline"`
	SummaryPolyline string `json:"summary_polyline"`
}

type Direction struct {
	Distance int64  `json:"distance"`
	Action   int    `json:"action"`
	Name     string `json:"name"`
}

type Route struct {
	Id            int64                    `json:"id"`
	Athlete       AthleteSummary           `json:"athlete"`
	Description   string                   `json:"description"`
	Distance      int                      `json:"distance"`
	ElevationGain int                      `json:"elevation_gain"`
	Map           Map                      `json:"map"`
	Name          string                   `json:"name"`
	Private       bool                     `json:"private"`
	Starred       bool                     `json:"starred"`
	Timestamp     int64                    `json:"timestamp"`
	RouteType     int                      `json:"type"`
	SubType       int                      `json:"sub_type"`
	Segments      []PersonalSegmentSummary `json:"segments"`
	Directions    []Direction              `json:"directions"`
}

type RoutesService struct {
	client *Client
}

func NewRoutesService(client *Client) *RoutesService {
	return &RoutesService{client}
}

/*********************************************************/

type RoutesGetCall struct {
	service *RoutesService
	id      int64
	ops     map[string]interface{}
}

func (s *RoutesService) Get(routeId int64) *RoutesGetCall {
	return &RoutesGetCall{
		service: s,
		id:      routeId,
		ops:     make(map[string]interface{}),
	}
}

func (c *RoutesGetCall) Do() (*Route, error) {
	data, err := c.service.client.run("GET", fmt.Sprintf("/routes/%d", c.id), c.ops)
	if err != nil {
		return nil, err
	}

	var route Route
	err = json.Unmarshal(data, &route)
	if err != nil {
		return nil, err
	}

	return &route, nil
}

/*********************************************************/
