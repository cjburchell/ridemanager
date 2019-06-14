package data

import "github.com/cjburchell/ridemanager/service/data/models"

type IService interface {
	GetUsers() ([]models.User, error)
	GetUser(athleteId models.AthleteId) (*models.User, error)
	GetStravaUser(athleteId int64) (*models.User, error)

	AddUser(user *models.User) error
	UpdateUser(user models.User) error
	DeleteUser(athleteId string) error

	GetOwnedActivities(ownerId models.AthleteId) ([]models.Activity, error)
	GetAthleteActivities(athleteId models.AthleteId) ([]models.Activity, error)
	GetAthleteActivitiesByState(athleteId models.AthleteId, state models.ActivityState) ([]models.Activity, error)
	GetAthletePrivateActivities(athleteId models.AthleteId) ([]models.Activity, error)

	GetActivityParticipants(activityId models.ActivityId) ([]models.Participant, error)
	AddParticipant(participant models.Participant) error
	UpdateParticipant(participant models.Participant) error
	DeleteParticipant(activityId models.ActivityId, athleteId string) error

	GetActivityParticipantsCount(activityId models.ActivityId) (int, error)
	GetActivityComments(activityId models.ActivityId) ([]models.Comment, error)
	AddComment(comment models.Comment) error

	GetPlaceCount(athleteId models.AthleteId, place int) (int, error)
	GetFinishedActivitiesCount(athleteId models.AthleteId) (int, error)

	GetActivitiesByPrivacy(activityPrivacy models.ActivityPrivacy) ([]models.Activity, error)
	GetActivitiesByState(athleteId models.AthleteId, state models.ActivityState) ([]models.Activity, error)

	AddActivity(user *models.Activity) error
	UpdateActivity(activity models.Activity) error
	DeleteActivity(activityId models.ActivityId) error
}
