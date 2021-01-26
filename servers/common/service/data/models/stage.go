package models

import "github.com/cjburchell/strava-go"

type Stage struct {
	SegmentId    int64               `json:"segment_id" bson:"segment_id"`
	Distance     float64             `json:"distance" bson:"distance"`
	Number       int                 `json:"number" bson:"number"`
	ActivityType strava.ActivityType `json:"activity_type" bson:"activity_type"`
	Name         string              `json:"name" bson:"name"`
	Map          []Point             `json:"map" bson:"map"`
	StartLatlng  strava.LatLng       `json:"start_latlng,omitempty" bson:"start_latlng"`
	EndLatlng    strava.LatLng       `json:"end_latlng,omitempty" bson:"end_latlng"`
}
