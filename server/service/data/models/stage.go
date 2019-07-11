package models

type Stage struct {
	SegmentId int `json:"segment_id" bson:"segment_id"`
	Distance  int `json:"distance" bson:"distance"`
	Number int `json:"number" bson:"number"`
	ActivityType ActivityType `json:"activity_type" bson:"activity_type"`
	Name string `json:"name" bson:"name"`
}
