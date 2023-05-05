package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Item struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:"title,omitempty"`
	Description string             `bson:"description,omitempty"`
	Country     string             `bson:"country,omitempty"`
	State       string             `bson:"state,omitempty"`
	City        string             `bson:"city,omitempty"`
	Address     string             `bson:"address,omitempty"`
	Photos      []string           `bson:"photos,omitempty"`
	UserId      int                `bson:"user,omitempty"`
}

type Items []Item
