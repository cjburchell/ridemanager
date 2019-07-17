package models

import (
	"time"
)

type Participant struct {
	Athlete        Athlete       `json:"athlete" bson:"athlete"`
	CategoryId     CategoryId    `json:"category_id" bson:"category_id"`
	Results        []Result      `json:"results" bson:"results"`
	Time           time.Time     `json:"time" bson:"time"`
	Rank           int           `json:"rank" bson:"rank"`
	OutOf          int           `json:"out_of" bson:"out_of"`
	StagesComplete int           `json:"stages_complete" bson:"stages_complete"`
}
