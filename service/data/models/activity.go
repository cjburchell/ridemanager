package models

import "time"

type ActivityType string

const (
	GroupRideActivity ActivityType = "group_ride"
	RaceActivity      ActivityType = "race"
	TriathlonActivity ActivityType = "triathlon"
)

type ActivityPrivacy string

const (
	PublicActivity  ActivityPrivacy = "public"
	PrivateActivity ActivityPrivacy = "private"
)

type ActivityState string

const (
	UpcommingActivity  ActivityState = "upcoming"
	InProgressActivity ActivityState = "in_progress"
	FinishedActivity   ActivityState = "finished"
)

type ActivityId string

type Activity struct {
	ActivityId      ActivityId `json:"activity_id" bson:"activity_id"`
	ActivityType    ActivityType
	OwnerId         AthleteId `json:"owner_id" bson:"owner_id"`
	Name            string
	Description     string
	StartTime       time.Time
	EndTime         time.Time
	TotalDistance   int
	Duration        float64
	TimeLeft        float64
	StartsIn        float64
	Privacy         ActivityPrivacy `json:"privacy" bson:"privacy"`
	Categories      []Category
	Stages          []Stage
	State           ActivityState `json:"state" bson:"state"`
	MaxParticipants int
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
		activity.State = FinishedActivity
	} else if activity.StartsIn <= 0 {
		activity.State = InProgressActivity
	} else {
		activity.State = UpcommingActivity
	}
}
