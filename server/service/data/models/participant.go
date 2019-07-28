package models

import (
	strava "github.com/cjburchell/go.strava"
	"github.com/cjburchell/ridemanager/service/stravaService"
	"time"
)

type Participant struct {
	Athlete        Athlete    `json:"athlete" bson:"athlete"`
	CategoryId     CategoryId `json:"category_id" bson:"category_id"`
	Results        []Result   `json:"results" bson:"results"`
	Time           time.Time  `json:"time" bson:"time"`
	Rank           int        `json:"rank" bson:"rank"`
	OutOf          int        `json:"out_of" bson:"out_of"`
	StagesComplete int        `json:"stages" bson:"stages"`
	OffsetTime     float64    `json:"offset_time" bson:"offset_time"`
}

func (participant *Participant) UpdateParticipantsResults(activity *Activity) error {
	participant.Results = make([]Result, len(activity.Stages))


	stravaService := stravaService.NewService("")

	for index, stage := range activity.Stages {
		participant.Results[index].SegmentId = stage.SegmentId
		participant.Results[index].StageNumber = stage.Number

		efforts, err := stravaService.SegmentsListEfforts(stage.SegmentId, participant.Athlete.StravaAthleteId, activity.StartTime, activity.EndTime)
		if err != nil{
			return err
		}

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
	}

	return nil
}
