package models

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
