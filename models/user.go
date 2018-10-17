package models

import (
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Uid      	bson.ObjectId 	`bson:"_id" json:"id"`
	Username 	string        	`bson:"username" json:"username"`
	Wokspaces	[]Workspace 	`bson: "wslist" json:"wslist"`
}
