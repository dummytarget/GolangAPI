package models

import "gopkg.in/mgo.v2/bson"

type (
	Account struct {
		Id      bson.ObjectId  `json:"id" bson:"_id,omitempty"`
		Name	string         `json:"name" bson:"name"`
		Score	float64        `json:"score" bson:"score"`
		Email   string         `json:"email" bson:"email"`
	}
)
