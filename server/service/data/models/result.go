package models

type Result struct {
	SegmentId   int64   `json:"segment_id" bson:"segment_id"`
	Time        float64 `json:"time" bson:"time"`
	Rank        int     `json:"rank" bson:"rank"`
	ActivityId  int     `json:"activity_id" bson:"activity_id"`
	StageNumber int     `json:"stage_number" json:"stage_number"`
}
