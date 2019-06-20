package models

type Stage struct {
	SegmentId string `json:"segment_id" bson:"segment_id"`
	Distance  int    `json:"distance" bson:"distance"`
}
