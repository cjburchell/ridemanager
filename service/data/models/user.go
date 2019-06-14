package models

import "github.com/cjburchell/go.strava"

const (
	UserRole  Role = "user"
	AdminRole Role = "admin"
)

type Role string

type AthleteId string

type User struct {
	Id                  AthleteId     `json:"id" bson:"id"`
	StravaAthleteId     int64         `json:"strava_athlete_id" bson:"strava_athlete_id"`
	StravaToken         string        `json:"strava_token" bson:"strava_token"`
	Role                Role          `json:"role" bson:"role"`
	MaxActiveActivities int           `json:"max_active_activities" bson:"max_active_activities"`
	Name                string        `json:"name" bson:"name"`
	Gender              strava.Gender `json:"sex" bson:"sex"`
	ProfileImage        string        `json:"profile" bson:"profile"`
	ProfileMediumImage  string        `json:"profile_medium" bson:"profile_medium"`
}

func NewUser(stravaAthleteId int64) *User {
	return &User{StravaAthleteId: stravaAthleteId, Role: UserRole, MaxActiveActivities: 3}
}
