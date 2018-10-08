package models

import (
	"gopkg.in/mgo.v2/bson"
)
// Represents a movie, we uses bson keyword to tell the mgo driver how to name
// the properties in mongodb document
type User struct {
	Uid      	bson.ObjectId 	`bson:"_id" json:"id"`
	username 	string        	`bson:"username" json:"username"`
	Wokspaces	[]Workspace 	`bson: "wslist" json:"wslist"`
}
