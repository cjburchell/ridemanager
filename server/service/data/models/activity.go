package models

import (
	"github.com/cjburchell/go.strava"
	"time"
)

type ActivityType string

var  ActivityTypes = struct {
	GroupRide ActivityType
	Race      ActivityType
	Triathlon ActivityType
}{"group_ride", "race", "triathlon"}

type ActivityPrivacy string

var Privacy = struct {
	Public ActivityPrivacy
	Private ActivityPrivacy
}{"public", "private"}

type ActivityState string

var ActivityStates = struct {
	Upcoming ActivityState
	InProgress ActivityState
	Finished   ActivityState
}{"upcoming", "in_progress", "finished"}

type Route struct {
	RouteId  int     `json:"id" bson:"id"`
	Name     string  `json:"name" bson:"name"`
	Distance float64 `json:"distance" bson:"distance"`
	Map      Map     `json:"map" bson:"map"`
}

type Map struct {
	Id              string `json:"id" bson:"id"`
	Polyline        string `json:"polyline" bson:"polyline"`
	SummaryPolyline string `json:"summary_polyline" bson:"summary_polyline"`
}

type Athlete struct {
	Id                 AthleteId     `json:"id" bson:"id"`
	StravaAthleteId    int64         `json:"strava_athlete_id" bson:"strava_athlete_id"`
	Name               string        `json:"name" bson:"name"`
	ProfileImage       string        `json:"profile" bson:"profile"`
	ProfileMediumImage string        `json:"profile_medium" bson:"profile_medium"`
	Gender             strava.Gender `json:"sex" bson:"sex"`
}

type ActivityId string

type Activity struct {
	ActivityId      ActivityId      `json:"activity_id" bson:"activity_id"`
	ActivityType    ActivityType    `json:"activity_type" bson:"activity_type"`
	Owner           Athlete         `json:"owner" bson:"owner"`
	Name            string          `json:"name" bson:"name"`
	Description     string          `json:"description" bson:"description"`
	StartTime       time.Time       `json:"start_time" bson:"start_time"`
	EndTime         time.Time       `json:"end_time" bson:"end_time"`
	TotalDistance   float64         `json:"total_distance" bson:"total_distance"`
	Duration        float64         `json:"duration" bson:"duration"`
	TimeLeft        float64         `json:"time_left" bson:"time_left"`
	StartsIn        float64         `json:"starts_in" bson:"starts_in"`
	Route           *Route          `json:"route" bson:"route"`
	Privacy         ActivityPrivacy `json:"privacy" bson:"privacy"`
	Categories      []Category      `json:"categories" bson:"categories"`
	Stages          []Stage         `json:"stages" bson:"stages"`
	Participants    []Participant   `json:"participants" bson:"participants"`
	State           ActivityState   `json:"state" bson:"state"`
	MaxParticipants int             `json:"max_participants" bson:"max_participants"`
}

func (activity *Activity) UpdateActivityState() {
	activity.TotalDistance = 0
	for _, item := range activity.Stages {
		activity.TotalDistance += item.Distance
	}

	activity.StartsIn = activity.StartTime.Sub(time.Now()).Seconds()
	activity.Duration = activity.EndTime.Sub(activity.EndTime).Seconds()
	activity.TimeLeft = activity.EndTime.Sub(time.Now()).Seconds()
	if activity.TimeLeft <= 0 {
		activity.State = ActivityStates.Finished
	} else if activity.StartsIn <= 0 {
		activity.State = ActivityStates.InProgress
	} else {
		activity.State = ActivityStates.Upcoming
	}
}
