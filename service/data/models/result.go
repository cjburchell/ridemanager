package models

import "time"

type Result struct {
	SegmentId  string     `json:"segmentId"`
	Time       time.Time  `json:"time"`
	ActivityId ActivityId `json:"activityId"`
}
