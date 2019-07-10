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
		u1 := uuid.NewV4()
		user.Id = models.AthleteId(u1.String())
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
	panic("implement me")
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
	err = s.db.C(activityCollection).Find(bson.M{"activity_id": bson.M{"$in": activityIds}, "privacy": models.PrivateActivity}).All(&activities)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return activities, nil
}

func (s service) GetActivityParticipantsCount(activityId models.ActivityId) (int, error) {
	var activity models.Activity
	err := s.db.C(usersCollection).Find(bson.M{"activity_id": activityId}).One(&activity)
	if err != nil {
		return 0, errors.WithStack(err)
	}

	return len(activity.Participants), nil
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
	var activities []models.Activity
	err := s.db.C(activityCollection).Find(bson.M{"privacy": activityPrivacy}).All(&activities)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return activities, nil
}

func (s service) GetActivitiesByState(state models.ActivityState) ([]models.Activity, error) {
	var activities []models.Activity
	err := s.db.C(activityCollection).Find(bson.M{"state": state}).All(&activities)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return activities, nil
}

func (s service) AddActivity(activity *models.Activity) error {
	if activity.ActivityId == "" {
		u1 := uuid.Must(uuid.NewV4(), errors.New("unable to create uuid"))
		activity.ActivityId = models.ActivityId(u1.String())
	}

	err := s.db.C(activityCollection).Insert(activity)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (s service) UpdateActivity(activity models.Activity) error {
	err := s.db.C(usersCollection).Update(bson.M{"activity_id": activity.ActivityId}, activity)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (s service) DeleteActivity(activityId models.ActivityId) error {
	err := s.db.C(usersCollection).Remove(bson.M{"activity_id": activityId})
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (s service) GetActivity(activityId models.ActivityId) (*models.Activity, error) {
	var activity models.Activity
	err := s.db.C(activityCollection).Find(bson.M{"activity_id": activityId}).One(&activity)

	if err == mgo.ErrNotFound {
		return nil, nil
	}

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &activity, err
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
