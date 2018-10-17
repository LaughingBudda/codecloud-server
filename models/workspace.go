package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Workspace struct {
	Wsid      bson.ObjectId `bson:"_id" json:"id"`
	Workspace_name string        `bson:"wsname" json:"wsname"`
}
