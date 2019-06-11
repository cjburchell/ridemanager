package models

import "time"

type Comment struct {
	text       string
	athleteId  AthleteId
	activityId ActivityId
	time       time.Time
}
