package data

import (
	"github.com/cjburchell/ridemanager/service/data/models"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type service struct {
	IService

	session *mgo.Session
	db      *mgo.Database
}

const dbName = "RideManager"
const usersCollection = "Users"
const activityCollection = "Activities"
const participantCollection = "Participants"

func (s service) GetUsers() ([]models.User, error) {
	var users []models.User
	err := s.db.C(usersCollection).Find(nil).All(&users)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return users, nil
}
func (s service) GetStravaUser(athleteId int64) (*models.User, error) {
	var user models.User
	err := s.db.C(usersCollection).Find(bson.M{"strava_athlete_id": athleteId}).One(&user)

	if err == mgo.ErrNotFound {
		return nil, nil
	}

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &user, err
}

func (s service) GetUser(athleteId models.AthleteId) (*models.User, error) {
	var user models.User
	err := s.db.C(usersCollection).Find(bson.M{"id": athleteId}).One(&user)

	if err == mgo.ErrNotFound {
		return nil, nil
	}

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &user, err
}

func (s service) AddUser(user *models.User) error {
	if user.Id == "" {
		user.Id = models.AthleteId(uuid.Must(uuid.NewV4()).String())
	}

	err := s.db.C(usersCollection).Insert(user)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (s service) UpdateUser(user models.User) error {
	err := s.db.C(usersCollection).Update(bson.M{"id": user.Id}, user)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (s service) DeleteUser(athleteId string) error {
	err := s.db.C(usersCollection).Remove(bson.M{"id": athleteId})
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (s service) GetOwnedActivities(ownerId models.AthleteId) ([]models.Activity, error) {
	var activities []models.Activity
	err := s.db.C(activityCollection).Find(bson.M{"owner_id": ownerId}).All(&activities)

	if err == mgo.ErrNotFound {
		return nil, nil
	}

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return activities, nil
}

func (s service) getParticipantActivities(athleteId models.AthleteId) ([]string, error) {
	var results []struct {
		Id string `bson:"activity_id"`
	}

	err := s.db.C(participantCollection).Find(bson.M{"athlete_id": athleteId}).All(&results)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	var activityIds = make([]string, len(results))
	for index, result := range results {
		activityIds[index] = result.Id
	}

	return activityIds, nil
}

func (s service) GetAthleteActivities(athleteId models.AthleteId) ([]models.Activity, error) {
	activityIds, err := s.getParticipantActivities(athleteId)
	if err != nil {
		return nil, err
	}

	var activities []models.Activity
	err = s.db.C(activityCollection).Find(bson.M{"activity_id": bson.M{"$in": activityIds}}).All(&activities)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return activities, nil
}

func (s service) GetAthleteActivitiesByState(athleteId models.AthleteId, state models.ActivityState) ([]models.Activity, error) {
	activityIds, err := s.getParticipantActivities(athleteId)
	if err != nil {
		return nil, err
	}

	var activities []models.Activity

	err = s.db.C(activityCollection).Find(bson.M{"activity_id": bson.M{"$in": activityIds}, "state": state}).All(&activities)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return activities, nil
}

func (s service) GetAthletePrivateActivities(athleteId models.AthleteId) ([]models.Activity, error) {

	activityIds, err := s.getParticipantActivities(athleteId)
	if err != nil {
		return nil, err
	}

	var activities []models.Activity
	err = s.db.C(activityCollection).Find(bson.M{"activity_id": bson.M{"$in": activityIds}, "Privacy": models.PrivateActivity}).All(&activities)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return activities, nil
}

func (s service) GetActivityParticipants(activityId models.ActivityId) ([]models.Participant, error) {
	var results []models.Participant
	err := s.db.C(participantCollection).Find(bson.M{"activity_id": activityId}).All(&results)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return results, nil
}

func (s service) AddParticipant(participant models.Participant) error {
	err := s.db.C(participantCollection).Insert(participant)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (s service) UpdateParticipant(participant models.Participant) error {
	err := s.db.C(participantCollection).Update(bson.M{"athlete_id": participant.AthleteId, "activity_id": participant.ActivityId}, participant)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (s service) DeleteParticipant(activityId models.ActivityId, athleteId string) error {
	err := s.db.C(participantCollection).Remove(bson.M{"athlete_id": athleteId, "activity_id": activityId})
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (s service) GetActivityParticipantsCount(activityId models.ActivityId) (int, error) {
	result, err := s.db.C(participantCollection).Find(bson.M{"activity_id": activityId}).Count()
	if err != nil {
		return 0, errors.WithStack(err)
	}

	return result, nil
}

func (s service) GetActivityComments(activityId models.ActivityId) ([]models.Comment, error) {
	panic("implement me")
}

func (s service) AddComment(comment models.Comment) error {
	panic("implement me")
}

func (s service) GetPlaceCount(athleteId models.AthleteId, place int) (int, error) {
	panic("implement me")
}

func (s service) GetFinishedActivitiesCount(athleteId models.AthleteId) (int, error) {
	panic("implement me")
}

func (s service) GetActivitiesByPrivacy(activityPrivacy models.ActivityPrivacy) ([]models.Activity, error) {
	panic("implement me")
}

func (s service) GetActivitiesByState(athleteId models.AthleteId, state models.ActivityState) ([]models.Activity, error) {
	panic("implement me")
}

func (s service) AddActivity(activity *models.Activity) error {
	panic("implement me")
}

func (s service) UpdateActivity(activity models.Activity) error {
	panic("implement me")
}

func (service) DeleteActivity(activityId models.ActivityId) error {
	panic("implement me")
}

func (s *service) setup(address string) error {
	session, err := mgo.Dial(address)
	if err != nil {
		return errors.WithStack(err)
	}

	err = session.Ping()
	if err != nil {
		return errors.WithStack(err)
	}

	s.db = session.DB(dbName)
	s.session = session

	return nil
}

func NewService(address string) (IService, error) {
	service := &service{}
	err := service.setup(address)
	return service, err
}
