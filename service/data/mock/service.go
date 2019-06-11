package mock

import (
	"github.com/cjburchell/ridemanager/service/data"
	"github.com/cjburchell/ridemanager/service/data/models"
)

type service struct {
}

func (service) GetUsers() ([]models.User, error) {
	panic("implement me")
}

func (service) GetUser(athleteId models.AthleteId) (models.User, error) {
	panic("implement me")
}

func (service) AddUser(user models.User) error {
	panic("implement me")
}

func (service) UpdateUser(user models.User) error {
	panic("implement me")
}

func (service) DeleteUser(athleteId string) error {
	panic("implement me")
}

func (service) GetOwnedActivities(ownerId models.AthleteId) ([]models.Activity, error) {
	panic("implement me")
}

func (service) GetAthleteActivities(athleteId models.AthleteId) ([]models.Activity, error) {
	panic("implement me")
}

func (service) GetAthleteActivitiesByState(athleteId models.AthleteId, state models.ActivityState) ([]models.Activity, error) {
	panic("implement me")
}

func (service) GetAthletePrivateActivities(athleteId models.AthleteId) ([]models.Activity, error) {
	panic("implement me")
}

func (service) GetActivityParticipants() ([]models.Participant, error) {
	panic("implement me")
}

func (service) AddParticipant(participant models.Participant) error {
	panic("implement me")
}

func (service) UpdateParticipant(participant models.Participant) error {
	panic("implement me")
}

func (service) DeleteParticipant(activityId models.ActivityId, athleteId string) error {
	panic("implement me")
}

func (service) GetActivityParticipantsCount(activityId models.ActivityId) (int, error) {
	panic("implement me")
}

func (service) GetActivityComments(activityId models.ActivityId) ([]models.Comment, error) {
	panic("implement me")
}

func (service) AddComment(comment models.Comment) error {
	panic("implement me")
}

func (service) GetPlaceCount(athleteId models.AthleteId, place int) (int, error) {
	panic("implement me")
}

func (service) GetFinishedActivitiesCount(athleteId models.AthleteId) (int, error) {
	panic("implement me")
}

func (service) GetActivitiesByPrivacy(activityPrivacy models.ActivityPrivacy) ([]models.Activity, error) {
	panic("implement me")
}

func (service) GetActivitiesByState(athleteId models.AthleteId, state models.ActivityState) ([]models.Activity, error) {
	panic("implement me")
}

func (service) AddActivity(user models.User) error {
	panic("implement me")
}

func (service) UpdateActivity(activity models.Activity) error {
	panic("implement me")
}

func (service) DeleteActivity(activityId models.ActivityId) error {
	panic("implement me")
}

func NewService() data.IService {
	return &service{}
}
