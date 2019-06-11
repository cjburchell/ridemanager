package models

import "time"

type Sex string

const (
	SexM Sex = "M"
	SexF Sex = "F"
)

type Participant struct {
	AthleteId      AthleteId  `json:"athlete_id" bson:"athlete_id"`
	ActivityId     ActivityId `json:"activity_id" bson:"activity_id"`
	ActivityState  ActivityState
	CategoryId     CategoryId
	Results        []Result
	Name           string
	Sex            Sex
	Time           time.Time
	Rank           int
	OutOf          int
	StagesComplete int
}
