package models

import "time"

type Result struct {
	SegmentId string    `json:"segment_id" bson:"segment_id"`
	Time      time.Time `json:"time" bson:"time"`
}
