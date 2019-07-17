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
	err := s.db.C(usersCollection).Find(bson.M{"athlete.id": athleteId}).One(&user)

	if err == mgo.ErrNotFound {
		return nil, nil
	}

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &user, err
}

func (s service) AddUser(user *models.User) error {
	if user.Athlete.Id == "" {
		u1 := uuid.NewV4()
		user.Athlete.Id = models.AthleteId(u1.String())
	}

	err := s.db.C(usersCollection).Insert(user)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (s service) UpdateUser(user models.User) error {
	err := s.db.C(usersCollection).Update(bson.M{"athlete.id": user.Athlete.Id}, user)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (s service) DeleteUser(athleteId string) error {
	err := s.db.C(usersCollection).Remove(bson.M{"athlete.id": athleteId})
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (s service) GetOwnedActivities(ownerId models.AthleteId) ([]models.Activity, error) {
	var activities []models.Activity
	err := s.db.C(activityCollection).Find(bson.M{"owner.id": ownerId}).All(&activities)

	if err == mgo.ErrNotFound {
		return nil, nil
	}

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return activities, nil
}

func (s service) GetAthleteActivities(athleteId models.AthleteId) ([]models.Activity, error) {
	var activities []models.Activity
	err := s.db.C(activityCollection).Find(bson.M{"participants": bson.M{"$elemMatch": bson.M{"athlete.id": athleteId}}}).All(&activities)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return activities, nil
}

func (s service) GetAthleteActivitiesByState(athleteId models.AthleteId, state models.ActivityState) ([]models.Activity, error) {
	var activities []models.Activity

	err := s.db.C(activityCollection).Find(bson.M{"participants": bson.M{"$elemMatch": bson.M{"athlete.id": athleteId}}, "state": state}).All(&activities)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return activities, nil
}

func (s service) GetAthleteActivitiesByStateCount(athleteId models.AthleteId, state models.ActivityState) (int, error) {
	count, err := s.db.C(activityCollection).Find(bson.M{"participants": bson.M{"$elemMatch": bson.M{"athlete.id": athleteId}}, "state": state}).Count()
	if err != nil {
		return 0, errors.WithStack(err)
	}

	return count, nil
}

func (s service) GetAthleteActivitiesPlaceCount(athleteId models.AthleteId, place int) (int, error) {
	count, err := s.db.C(activityCollection).Find(bson.M{"participants": bson.M{"$elemMatch": bson.M{"athlete.id": athleteId, "rank": place}}, "state": models.ActivityStates.Finished}).Count()
	if err != nil {
		return 0, errors.WithStack(err)
	}

	return count, nil
}

func (s service) GetAthletePrivateActivities(athleteId models.AthleteId) ([]models.Activity, error) {

	var activities []models.Activity
	err := s.db.C(activityCollection).Find(bson.M{"participants": bson.M{"$elemMatch": bson.M{"athlete.id": athleteId}}, "privacy": models.Privacy.Private}).All(&activities)
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

func (s service) AddActivity(activity *models.Activity) (models.ActivityId, error) {
	if activity.ActivityId == "" {
		u1 := uuid.NewV4()
		activity.ActivityId = models.ActivityId(u1.String())
	}

	err := s.db.C(activityCollection).Insert(activity)
	if err != nil {
		return models.ActivityId(""), errors.WithStack(err)
	}
	return activity.ActivityId, nil
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
