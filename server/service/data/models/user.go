package models

const (
	UserRole  Role = "user"
	AdminRole Role = "admin"
)

type Role string

type AthleteId string

type Token struct {
	AccessToken  string `bson:"access_token" json:"access_token"`
	RefreshToken string `bson:"refresh_token" json:"refresh_token"`
	ExpiresAt    int64  `bson:"expires_at" json:"expires_at"`
}

type User struct {
	Athlete             Athlete `json:"athlete" bson:"athlete"`
	StravaToken         Token   `json:"strava_token" bson:"strava_token"`
	Role                Role    `json:"role" bson:"role"`
	MaxActiveActivities int     `json:"max_active_activities" bson:"max_active_activities"`
}

func NewUser(stravaAthleteId int32) *User {
	return &User{Athlete: Athlete{StravaAthleteId: stravaAthleteId}, Role: UserRole, MaxActiveActivities: 3}
}
