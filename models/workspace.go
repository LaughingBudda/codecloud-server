package models

import "gopkg.in/mgo.v2/bson"

// Represents a movie, we uses bson keyword to tell the mgo driver how to name
// the properties in mongodb document
type Workspace struct {
	Wsid      bson.ObjectId `bson:"_id" json:"id"`
	Workspace_name string        `bson:"wsname" json:"wsname"`
}
