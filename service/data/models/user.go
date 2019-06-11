package models

const (
	UserRole  Role = "user"
	AdminRole Role = "admin"
)

type Role string

type AthleteId string

type User struct {
	Id                  AthleteId `json:"id" bson:"id"`
	StravaAthleteId     string    `json:"strava_athlete_id" bson:"strava_athlete_id"`
	Role                Role      `json:"role" bson:"role"`
	MaxActiveActivities int       `json:"max_active_activities" bson:"max_active_activities"`
	Name                string    `json:"name" bson:"name"`
	Email               string    `json:"email" bson:"email"`
}
