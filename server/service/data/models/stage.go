package models

import "github.com/cjburchell/go.strava"

type Stage struct {
	SegmentId     int64               `json:"segment_id" bson:"segment_id"`
	Distance      float64             `json:"distance" bson:"distance"`
	Number        int                 `json:"number" bson:"number"`
	ActivityType  strava.ActivityType `json:"activity_type" bson:"activity_type"`
	Name          string              `json:"name" bson:"name"`
	Map           Map                 `json:"map" bson:"map"`
	StartLocation strava.Location     `json:"start_latlng" bson:"start_latlng"`
	EndLocation   strava.Location     `json:"end_latlng" bson:"end_latlng"`
}
