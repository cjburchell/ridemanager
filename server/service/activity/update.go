package activity

import (
	"github.com/cjburchell/ridemanager/service/data"
	"github.com/cjburchell/ridemanager/service/data/models"
)

func UpdateAll(activities []*models.Activity, service data.IService) error {
	for _, activity := range activities{
		changed := activity.UpdateState()

		if changed{
			err := service.UpdateActivity(*activity)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

