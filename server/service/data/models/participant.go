package models

import (
	"github.com/cjburchell/go.strava"
	"github.com/cjburchell/ridemanager/service/stravaService"
	"time"
)

type Participant struct {
	Athlete        Athlete    `json:"athlete" bson:"athlete"`
	CategoryId     CategoryId `json:"category_id" bson:"category_id"`
	Results        []Result   `json:"results" bson:"results"`
	Time           float64    `json:"time" bson:"time"`
	Rank           int        `json:"rank" bson:"rank"`
	OutOf          int        `json:"out_of" bson:"out_of"`
	StagesComplete int        `json:"stages" bson:"stages"`
	OffsetTime     float64    `json:"offset_time" bson:"offset_time"`
}

func filterSegments(ss []*strava.SegmentEffortSummary, test func(*strava.SegmentEffortSummary) bool) (ret []*strava.SegmentEffortSummary) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}

func filterResults(ss []Result, test func(Result) bool) (ret []Result) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}

func filterParticipants(ss []*Participant, test func(*Participant) bool) (ret []*Participant) {
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

func (participant *Participant) UpdateParticipantsResults(activity *Activity, accessToken string) error {
	participant.Results = make([]Result, len(activity.Stages))


	ss := stravaService.NewService(accessToken)

	for index, stage := range activity.Stages {
		participant.Results[index].SegmentId = stage.SegmentId
		participant.Results[index].StageNumber = stage.Number

		efforts, err := ss.SegmentsListEfforts(stage.SegmentId, participant.Athlete.StravaAthleteId, activity.StartTime, activity.EndTime)
		if err != nil{
			return err
		}

		efforts = filterSegments(efforts, func(summary *strava.SegmentEffortSummary) bool {
			return inTimeSpan(activity.StartTime, activity.EndTime, summary.StartDate)
		})

		if len(efforts) == 0 {
			continue
		}

		var bestEffort *strava.SegmentEffortSummary = nil
		for _, effort := range efforts{
			if bestEffort == nil {
				bestEffort = effort
				continue
			}

			if bestEffort.ElapsedTime > effort.ElapsedTime{
				bestEffort = effort
			}
		}

		if bestEffort == nil {
			continue
		}

		participant.Results[index].Time = float64(bestEffort.ElapsedTime)
		participant.Results[index].ActivityId = bestEffort.Id
	}

	if activity.ActivityType == ActivityTypes.Triathlon || activity.ActivityType == ActivityTypes.Race{
		complete:= filterResults(participant.Results, func(result Result) bool {
			return result.ActivityId != 0
		})

		participant.Time = 0
		participant.StagesComplete = len(complete)
		if participant.StagesComplete != 0{
			for _,stage := range participant.Results{
				participant.Time += stage.Time
			}
		}
	}

	return nil
}
