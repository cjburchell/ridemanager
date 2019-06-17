package contract

import (
	"github.com/cjburchell/go.strava"
	"github.com/cjburchell/ridemanager/service/data/models"
)

type User struct {
	StravaAthleteId     int64         `json:"strava_athlete_id" `
	Role                models.Role   `json:"role"`
	MaxActiveActivities int           `json:"max_active_activities" `
	FirstName           string        `json:"first_name"`
	LastName            string        `json:"last_name" `
	Gender              strava.Gender `json:"sex" `
	ProfileImage        string        `json:"profile"`
	ProfileMediumImage  string        `json:"profile_medium"`
}

func NewUser(user models.User) *User {
	contract := &User{}
	contract.StravaAthleteId = user.StravaAthleteId
	contract.Role = user.Role
	contract.MaxActiveActivities = user.MaxActiveActivities
	contract.FirstName = user.FirstName
	contract.LastName = user.LastName
	contract.Gender = user.Gender
	contract.ProfileImage = user.ProfileImage
	contract.ProfileMediumImage = user.ProfileMediumImage
	return contract
}
