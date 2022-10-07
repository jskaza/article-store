package models

import (
	"html/template"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Paper struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Title     string             `bson:"title,omitempty"`
	Category  string             `bson:"category,omitempty"`
	Author    string             `bson:"author,omitempty"`
	Content   string             `bson:"content,omitempty"`
	Approvals int32              `bson:"approvals,omitempty"`
	Abstract  string             `bson:"abstract,omitempty"`
	HexID     string             `bson:"hexID,omitempty"`
}

type DisplayPaper struct {
	Title    string        `bson:"title,omitempty"`
	Category string        `bson:"category,omitempty"`
	Author   string        `bson:"author,omitempty"`
	Content  template.HTML `bson:"content,omitempty"`
	Abstract string        `bson:"abstract,omitempty"`
}

type User struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `bson:"name,omitempty"`
}
