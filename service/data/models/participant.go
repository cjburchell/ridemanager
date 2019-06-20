package models

import (
	"time"

	strava "github.com/cjburchell/go.strava"
)

type Participant struct {
	AthleteId      AthleteId     `json:"athlete_id" bson:"athlete_id"`
	CategoryId     CategoryId    `json:"category_id" bson:"category_id"`
	Results        []Result      `json:"results" bson:"results"`
	Name           string        `json:"name" bson:"name"`
	Sex            strava.Gender `json:"sex" bson:"sex"`
	Time           time.Time     `json:"time" bson:"time"`
	Rank           int           `json:"rank" bson:"rank"`
	OutOf          int           `json:"out_of" bson:"out_of"`
	StagesComplete int           `json:"stages_complete" bson:"stages_complete"`
}
