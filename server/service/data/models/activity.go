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
	UpcomingActivity   ActivityState = "upcoming"
	InProgressActivity ActivityState = "in_progress"
	FinishedActivity   ActivityState = "finished"
)

type ActivityId string

type Activity struct {
	ActivityId      ActivityId      `json:"activity_id" bson:"activity_id"`
	ActivityType    ActivityType    `json:"activity_type" bson:"activity_type"`
	OwnerId         AthleteId       `json:"owner_id" bson:"owner_id"`
	Name            string          `json:"name" bson:"name"`
	Description     string          `json:"description" bson:"description"`
	StartTime       time.Time       `json:"start_time" bson:"start_time"`
	EndTime         time.Time       `json:"end_time" bson:"end_time"`
	TotalDistance   int             `json:"total_distance" bson:"total_distance"`
	Duration        float64         `json:"duration" bson:"duration"`
	TimeLeft        float64         `json:"time_left" bson:"time_left"`
	StartsIn        float64         `json:"starts_in" bson:"starts_in"`
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
		activity.State = FinishedActivity
	} else if activity.StartsIn <= 0 {
		activity.State = InProgressActivity
	} else {
		activity.State = UpcomingActivity
	}
}
