package model

import (
	"gopkg.in/mgo.v2/bson"
)

type Ticker struct {
	Id_  bson.ObjectId `json:"_id" bson:"_id"`
	Name string        `json:"name" bson:"name"`
}
