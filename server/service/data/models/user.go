package models

const (
	UserRole  Role = "user"
	AdminRole Role = "admin"
)

type Role string

type AthleteId string

type User struct {
	Athlete             Athlete `json:"athlete" bson:"athlete"`
	StravaToken         string  `json:"strava_token" bson:"strava_token"`
	Role                Role    `json:"role" bson:"role"`
	MaxActiveActivities int     `json:"max_active_activities" bson:"max_active_activities"`
}

func NewUser(stravaAthleteId int64) *User {
	return &User{Athlete: Athlete{StravaAthleteId: stravaAthleteId}, Role: UserRole, MaxActiveActivities: 3}
}
