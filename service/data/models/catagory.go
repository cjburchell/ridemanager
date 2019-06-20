package models

type CategoryId string

type Category struct {
	CategoryId CategoryId `json:"category_id" json:"category_id"`
	Name       string     `json:"name" bson:"name"`
}
