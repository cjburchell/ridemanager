package activity

import (
	"github.com/cjburchell/ridemanager/common/service/data"
	"github.com/cjburchell/ridemanager/common/service/data/models"
	"github.com/cjburchell/ridemanager/common/service/results"
	"github.com/cjburchell/ridemanager/common/service/stravaService"
)

func UpdateAll(activities []*models.Activity, service data.IService, updateStanding bool, authenticator stravaService.Authenticator) error {
	for _, activity := range activities {
		changed := activity.UpdateState()

		if updateStanding && activity.State != models.ActivityStates.Upcoming {
			// only update the state of finished and in progress activities
			err := updateActivityStandings(activity, service, authenticator)
			if err != nil {
				return err
			}
			changed = true
		}

		if changed {
			err := service.UpdateActivity(*activity)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func updateActivityStandings(activity *models.Activity, service data.IService, authenticator stravaService.Authenticator) error {
	for _, participant := range activity.Participants {
		user, err := service.GetUser(participant.Athlete.Id)
		if err != nil {
			return err
		}

		err = results.UpdateParticipant(participant, activity, stravaService.GetTokenManager(authenticator, participant.Athlete.Id, service, &user.StravaToken))
		if err != nil {
			return err
		}
	}

	results.Update(activity)
	return nil
}
