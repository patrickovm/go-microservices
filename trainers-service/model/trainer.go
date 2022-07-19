package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Trainer struct {
	Id            primitive.ObjectID `bson:"_id,omitempty"`
	FirstName     string             `bson:"first_name,omitempty"`
	LastName      string             `bson:"last_name,omitempty"`
	Age           int32              `bson:"age,omitempty"`
	FinishedGames int32              `bson:"finished_games,omitempty"`
}
