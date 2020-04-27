package update

import (
	"sort"
	"time"

	"github.com/cjburchell/ridemanager/common/service/data/models"
	"github.com/cjburchell/ridemanager/common/service/stravaService"
	"github.com/cjburchell/strava-go"
)

func filterSegments(ss []strava.DetailedSegmentEffort, test func(strava.DetailedSegmentEffort) bool) (ret []strava.DetailedSegmentEffort) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}

func filterResults(ss []models.Result, test func(models.Result) bool) (ret []models.Result) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}

func filterParticipants(ss []*models.Participant, test func(*models.Participant) bool) (ret []*models.Participant) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}

func filterResultItem(ss []models.ResultItem, test func(models.ResultItem) bool) (ret []models.ResultItem) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}

func inTimeSpan(start, end, check time.Time) bool {
	return check.After(start) && check.Before(end)
}

func Standings(activity *models.Activity) {
	if activity.ActivityType == models.ActivityTypes.Triathlon || activity.ActivityType == models.ActivityTypes.Race {
		return
	}

	genders := []string{"M", "F"}
	for _, category := range activity.Categories {
		for _, gender := range genders {
			stageParticipants := filterParticipants(activity.Participants, func(participant *models.Participant) bool {
				return participant.CategoryId == category.CategoryId && participant.Athlete.Sex == gender
			})

			stageCount := len(activity.Stages)
			finishedParticipants := filterParticipants(stageParticipants, func(participant *models.Participant) bool {
				return participant.StagesComplete == stageCount
			})

			sort.Slice(finishedParticipants, func(i, j int) bool {
				return finishedParticipants[i].Time < finishedParticipants[j].Time
			})

			activity.CalculateRank(finishedParticipants, len(stageParticipants))

			if len(finishedParticipants) != 0 {
				topParticipant := finishedParticipants[0]
				activity.CalculateOffset(finishedParticipants, topParticipant)
			}

			for stageIndex := range activity.Stages {
				results := make([]models.ResultItem, len(stageParticipants))

				for index, participant := range stageParticipants {
					results[index].Participant = participant
					results[index].Result = &participant.Results[stageIndex]
				}

				sortedResults := filterResultItem(results, func(item models.ResultItem) bool {
					return item.Result.ActivityId != 0
				})

				sort.Slice(sortedResults, func(i, j int) bool {
					return sortedResults[i].Result.Time < sortedResults[j].Result.Time
				})

				stageRank := 0
				for _, item := range sortedResults {
					stageRank++
					item.Result.Rank = stageRank
				}
			}
		}
	}
}

func ParticipantsResults(participant *models.Participant, activity *models.Activity, accessToken stravaService.TokenManager) error {
	participant.Results = make([]models.Result, len(activity.Stages))

	ss := stravaService.NewService(accessToken)

	for index, stage := range activity.Stages {
		participant.Results[index].SegmentId = stage.SegmentId
		participant.Results[index].StageNumber = stage.Number

		efforts, err := ss.SegmentsListEfforts(stage.SegmentId, 0, 100)
		if err != nil {
			return err
		}

		efforts = filterSegments(efforts, func(summary strava.DetailedSegmentEffort) bool {
			return inTimeSpan(activity.StartTime, activity.EndTime, summary.StartDate)
		})

		if len(efforts) == 0 {
			continue
		}

		var bestEffort *strava.DetailedSegmentEffort = nil
		for _, effort := range efforts {
			if bestEffort == nil {
				bestEffort = &effort
				continue
			}

			if bestEffort.ElapsedTime > effort.ElapsedTime {
				bestEffort = &effort
			}
		}

		if bestEffort == nil {
			continue
		}

		participant.Results[index].Time = float64(bestEffort.ElapsedTime)
		participant.Results[index].ActivityId = bestEffort.Activity.Id
	}

	if activity.ActivityType == models.ActivityTypes.Triathlon || activity.ActivityType == models.ActivityTypes.Race {
		complete := filterResults(participant.Results, func(result models.Result) bool {
			return result.ActivityId != 0
		})

		participant.Time = 0
		participant.StagesComplete = len(complete)
		if participant.StagesComplete != 0 {
			for _, stage := range participant.Results {
				participant.Time += stage.Time
			}
		}
	}

	return nil
}
